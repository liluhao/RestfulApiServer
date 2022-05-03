package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zxmrlc/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"apiserver/config"
	"apiserver/model"
	v "apiserver/pkg/version"
	"apiserver/router"
	"apiserver/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ") //MarshalIndent类似Marshal但会使用缩进将输出格式化。
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// 初始化配置、初始化日志
	//读取配置文件方式是“从配置文件读取”
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//初始化数据库
	model.DB.Init()
	defer model.DB.Close() //最后要把数据库关闭

	// 设置gin的运行模式，gin有3种运行模式：debug、release、test，其中debug模式会打印很多信息
	gin.SetMode(viper.GetString("runmode"))

	//创造引擎
	g := gin.New()

	//加载路由
	router.Load(g, middleware.Logging(), middleware.RequestId())

	// Ping 服务器以确保路由器正常工作。
	go func() {
		if err := pingServer(); err != nil {
			//引用的是"github.com/zxmrlc/log"里的
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	/*
		1.Level DEBUG :指出细粒度信息事件对调试应用程序是非常有帮助的。
		2.levelINFO :表明 消息在粗粒度级别上突出强调应用程序的运行过程。
		3.Level WARN :表明会出现潜在错误的情形。
		4.Level ERROR: level指出虽然发生错误事件，但仍然不影响系统的继续运行
		5.Level FATAL:指出每个严重的错误事件将会导致应用程序的退出
	*/

	// Start to listening the incoming requests.
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	//url: http://127.0.0.1:8080   # pingServer函数请求的API服务器的ip:port
	// addr: :8081
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

//ping的次数由config.yaml里的max_ping_count所决定
func pingServer() error {
	//GetInt函数可以把传入的Int给转换成为Int类型
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health") //http里的函数，就会得到一个响应体
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.") //引用的是"github.com/zxmrlc/log"里的
		//因为go开启了协程，所以要保证ping够10次
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

/*
有时候AP!进程起来不代表API服务器正常，笔者曾经就遇到过这种问题：API进程存在，但是服务器却不能对
外提供服务。因此在启动API服务器时，如果能够最后做一个自检会更好些。笔者在apiserver中也添加了自检
程序，在启动HTTP端口前go一个pingServer协程，启动HTTP端口后，该协程不断地ping/sd/healt
路径，如果失败次数超过一定次数，则终止HTTP服务器进程。通过自检可以最大程度地保证启动后的API服
务器处于健康状态。
*/
