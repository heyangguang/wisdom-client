package collector

import (
	"fmt"
	"net"
	"time"
	"wisdom-client/wisdom-client/conf"
	"wisdom-client/wisdom-client/logger"
)

type ServiceStatus struct {
	Ip     string    `json:"ip"`
	Port   string    `json:"port"`
	Name   string    `json:"name"`
	Status bool      `json:"status"`
	Time   time.Time `json:"time"`
	Tag    string    `json:"tag"`
}

func NewServiceStatus(ip, port, name string) ServiceStatus {
	return ServiceStatus{
		Ip:     ip,
		Port:   port,
		Name:   name,
		Status: false,
		Time:   time.Time{},
		Tag:    "",
	}
}

// 循环探测 秒 单位
func (s *ServiceStatus) loopTcp(timeout, loopTime time.Duration, tag string) {
	for {
		s.tcpGather(timeout, tag)
		time.Sleep(loopTime * time.Second)
	}
}

// 核心探测功能
func (s *ServiceStatus) tcpGather(timeout time.Duration, tag string) {
	address := net.JoinHostPort(s.Ip, s.Port)
	// 超时 秒 单位
	if _, err := net.DialTimeout("tcp", address, timeout*time.Second); err != nil {
		s.Time = time.Now()
		s.Status = false
		s.Tag = tag
		if err := httpRequest(conf.YamlObj.Server.Host, conf.YamlObj.Server.Secret, s); err != nil {
			logger.Error("request http error, err: " + err.Error())
			return
		}
		logger.Debug(fmt.Sprintf("%v connect error, err: %s", s, err.Error()))
		logger.Info(fmt.Sprintf("%s connect error", s.Name))
	} else {
		//time.Sleep(timeout*time.Second)
		s.Time = time.Now()
		s.Status = true
		s.Tag = tag
		if err := httpRequest(conf.YamlObj.Server.Host, conf.YamlObj.Server.Secret, s); err != nil {
			logger.Error("request http error, err: " + err.Error())
			return
		}
		logger.Debug(fmt.Sprintf("%v connect success", s))
		logger.Info(fmt.Sprintf("%s connect success", s.Name))
	}
}
