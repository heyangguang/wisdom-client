package conf

import (
	"github.com/spf13/viper"
	"wisdom-client/wisdom-client/logger"
)

var (
	ViperConfig *viper.Viper
	YamlObj *YamlSetting
)


func InitConfig(path string)  {
	ViperConfig = viper.New()
	ViperConfig.SetConfigFile(path)
	ViperConfig.SetConfigType("yaml")
	if err := ViperConfig.ReadInConfig(); err != nil {
		logger.Error("read config error, err: " + err.Error())
	}
	PareYaml()
}

func PareYaml() {
	if err := ViperConfig.Unmarshal(&YamlObj); err != nil {
		logger.Error("pareYaml error, err: " + err.Error())
	}
}