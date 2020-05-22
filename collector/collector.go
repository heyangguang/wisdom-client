package collector

import (
	"time"
	wisdomClient "wisdom-client/wisdom-client"
	"wisdom-client/wisdom-client/conf"
)

// 处理探测服务列表

func StartMonitor() {
	go mysqlApp()
	go elasticSearchApp()
}


func mysqlApp() {
	app := conf.YamlObj.MySQLApplication
	registerApp(app.HostPool, app.Timeout, app.LoopTime, app.Tag)
}


func elasticSearchApp() {
	app := conf.YamlObj.ElasticSearchApplication
	registerApp(app.HostPool, app.Timeout, app.LoopTime, app.Tag)
}


// 服务注册
func registerApp(hostPoll []string, timeout, loopTime time.Duration, tag string) {
	for _, value := range hostPoll {
		strSlice := wisdomClient.StrSeparate(value)
		obj := NewServiceStatus(strSlice[0], strSlice[1], strSlice[2])
		go obj.loopTcp(timeout,loopTime, tag)
	}
}