package main

import (
	"fmt"
	"net/http"
	"wisdom-client/collector"
	wisdomClient "wisdom-client/wisdom-client"
	"wisdom-client/wisdom-client/conf"
	"wisdom-client/wisdom-client/logger"
)

func main()  {
	// 初始化日志模块
	logPath := wisdomClient.BaseDir() + "/logs/wisdom-client.log"
	err := logger.InitLogger(logPath, 1, 7, 10, "DEBUG")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 初始化配置文件
	configPath := wisdomClient.BaseDir() + "/conf/config.yaml"
	conf.InitConfig(configPath)


	// 启动核心
	collector.StartMonitor()


	// 启动服务
	fmt.Println("WisdomClient ListenAndServer 9090...")
	_ = http.ListenAndServe(":9090", nil)
}
