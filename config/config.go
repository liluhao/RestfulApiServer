package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
)

//Config结构体还有3个方法，分别是initConfig()、initLog()、watchConfig()
type Config struct {
	Name string
}

//解析并监控配置文件、初始化日志包;返回值设置为error
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	//
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	//这里要注意，日志初始化函数c.initLog()要放在配置初始化函数c.initConfig()之后，因为日志初始
	c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

//返回值设置为error
/*设置并解析配置文件。如果指定了配置文件*cfg不为空，则解析指定的配置文件，否则解析默认的配置
文件conf/config,yaml。通过指定配置文件可以很方便地连接不同的环境（开发环境、测试环境）并加
载不同的配置，方便开发和测试。*/
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件;参数填写路径
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML

	//通过如下设置可以使程序读取环境变量;
	//前缀与配置名称需要大写，二者用_连接，比如APISERVER_RUNMODE。如果配置项是嵌套的，情况可类推，比如APISERVER_DB_NAME
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),       //writers:输出位置，有两个可选项：file和stdout。选择fle会将日志记录到logger_file指定 的日志文件中，选择stdout会将日志输出到标准输出，当然也可以两者同时选择
		LoggerLevel:    viper.GetString("log.logger_level"),  //logger_level:日志级别，DEBUG、INFO、WARN、ERROR、FATAL
		LoggerFile:     viper.GetString("log.logger_file"),   //logger_fi1e:日志文件
		LogFormatText:  viper.GetBool("log.log_format_text"), //log format text:日志的输出格式，JSON或者plaintext,true会输出成非JSON格式， false 会输出成JSON格式
		RollingPolicy:  viper.GetString("log.rollingPolicy"), //rollingPolicy:rotate依据，可选的有daily和size。如果选daily则根据天进行转存，如果是size则根据大小进行转存
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),  //log_rotate_date:rotate转存时间，配合rollingPolicy:daily使用
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),  //log_rotate_size:rotate转存大小，配合rollingPolicy:size使用
		LogBackupCount: viper.GetInt("log.log_backup_count"), //1 og_backup_count：当日志文件达到转存标准时，Iog系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
	}

	log.InitWithConfig(&passLagerCfg)
}

// 监控配置文件变化并热加载程序;
//通过该函数的viper设置，可以使viper监控配置文件变更，如有变更则热更新程序。
//所谓热更新是指：可以不重启AP!进程，使API加载最新配置项的值。
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
