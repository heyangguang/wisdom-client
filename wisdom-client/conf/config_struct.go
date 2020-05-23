package conf

import "time"

type YamlSetting struct {
	Server                   Server
	MySQLApplication         MySQLApplication
	ElasticSearchApplication ElasticSearchApplication
}

type Server struct {
	Host   string
	Secret string
}

type MySQLApplication struct {
	LoopTime time.Duration
	Timeout  time.Duration
	HostPool []string
}

type ElasticSearchApplication struct {
	LoopTime time.Duration
	Timeout  time.Duration
	HostPool []string
}
