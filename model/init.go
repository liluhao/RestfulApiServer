package model

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//因为一个AP1服务器可能需要同时问多个数据库，为了对多个数据库进行初始化和连接管理，这里定义了一个 叫Database的struct:
type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

//Database结构体有个Init()方法用来初始化连接
/*Init()方法会调用GetSelfDB()和GetDockerDB()函数同时创建两个Database 的数据库对象。
这两个Get函数最终都会调用func openDB(username,.password,addr,name string)*gorm.DB函数
来建立数据库连接，不同数据库实例传入不同的username、password、addr和名字信息，从而建立不同的数据
库连接。*/
func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}
func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

func (db *Database) Close() {
	DB.Self.Close() //直接利用gorm的关闭函数
	DB.Docker.Close()
}
