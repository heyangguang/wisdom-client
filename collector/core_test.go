package collector

import (
	"fmt"
	"os"
	"testing"
	wisdomClient "wisdom-client/wisdom-client"
	"wisdom-client/wisdom-client/logger"
)


func TestTcpGather(t *testing.T)  {
	logPath := wisdomClient.BaseDir() + "wisdom-test.log"
	_ = logger.InitLogger(logPath, 1, 7, 10, "DEBUG")
	service := NewServiceStatus("127.0.0.1", "330346", "mysql")
	service.tcpGather(10, "test")
	fmt.Println(service)
	_ = os.Remove(logPath)
}
