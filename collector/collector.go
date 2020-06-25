package collector

import (
	"time"
	wisdomClient "wisdom-client/wisdom-client"
	"wisdom-client/wisdom-client/conf"
	"wisdom-client/wisdom-client/logger"
)

// 处理探测服务列表

func StartMonitor() {
	go mysqlApp()
	go elasticSearchApp()
	go kafkaApp()
	go KubernetesApp()
}

func mysqlApp() {
	logger.Info("Start MySQL Monitor Slope")
	app := conf.YamlObj.MySQLApplication
	registerApp(app.HostPool, app.Timeout, app.LoopTime, "MySQL")
}

func elasticSearchApp() {
	logger.Info("Start ElasticSearch Monitor Slope")
	app := conf.YamlObj.ElasticSearchApplication
	registerApp(app.HostPool, app.Timeout, app.LoopTime, "ElasticSearch")
}

func kafkaApp() {
	logger.Info("Start Kafka Monitor Slope")
	app := conf.YamlObj.KafkaApplication
	registerApp(app.HostPool, app.Timeout, app.LoopTime, "Kafka")
}

func KubernetesApp() {
	logger.Info("Start Kubernetes Monitor Slope")
	app := conf.YamlObj.KubernetesApplication
	registerApp(app.HostPool, app.Timeout, app.LoopTime, "Kubernetes")
}

// 服务注册
func registerApp(hostPoll []string, timeout, loopTime time.Duration, tag string) {
	for _, value := range hostPoll {
		strSlice := wisdomClient.StrSeparate(value)
		obj := NewServiceStatus(strSlice[0], strSlice[1], strSlice[2])
		go obj.loopTcp(timeout, loopTime, tag)
	}
}
