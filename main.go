package main

import (
	"flag"
	"fmt"

	"github.com/shfd/repo/application"
	"github.com/shfd/repo/config"
	log "github.com/sirupsen/logrus"
)

var (
	configFile = flag.String("c", config.DefaultConfigPath, "设置项目配置文件地址`<path>`，可选")
	version    = flag.Bool("v", false, "打印版本号并退出")
	command    string
)

// Version 程序版本号
const Version = "0.0.1"

func main() {
	// log.SetLevel(log.TraceLevel)
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "usage: %s [-options] [command]\n", os.Args[0])
	// 	fmt.Fprintf(os.Stderr, "Commands:\n")
	// 	fmt.Fprintf(os.Stderr, "  start 启动项目\n")
	// 	fmt.Fprintf(os.Stderr, "Options:\n")
	// 	flag.PrintDefaults()
	// }
	// flag.Parse()

	// log.Debugln(os.Args)

	// if (os.Args[len(os.Args)-1])[0] != '-' {
	// 	command = os.Args[len(os.Args)-1]
	// }

	// if *version {
	// 	fmt.Println(fmt.Sprintf("版本号：%s", Version))
	// 	return
	// }

	// switch command {
	// case "start":
	fmt.Println("server starting...")
	// default:
	// 	flag.Usage()
	// 	return
	// }

	// if(os.Args[-1])

	irisServer := application.NewServer(config.LoadConfig(*configFile))
	log.Info(config.GetConfig())
	if irisServer == nil {
		panic("http server 初始化失败")
	}

	// if libs.IsPortInUse(libs.Config.Port) {
	// 	if !irisServer.Status {
	// 		panic(fmt.Sprintf("端口 %d 已被使用\n", libs.Config.Port))
	// 	}
	// 	irisServer.Stop() // 停止
	// }

	err := irisServer.Start()
	if err != nil {
		panic(fmt.Sprintf("http server 启动失败: %+v", err))
	}

	// logging.InfoLogger.Infof("http server %s:%d start", libs.Config.Host, libs.Config.Port)

}
