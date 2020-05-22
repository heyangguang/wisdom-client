package conf

import (
	"path/filepath"
	"testing"
	wisdomClient "wisdom-client/wisdom-client"
)

func TestPareYaml(t *testing.T) {
	logPath := filepath.Dir(filepath.Dir(wisdomClient.BaseDir())) + "/conf/config.yaml"
	InitConfig(logPath)
}
