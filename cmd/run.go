/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"wisdom-client/collector"
	"wisdom-client/wisdom-client/conf"
	"wisdom-client/wisdom-client/logger"
)

var config string
var port int
var log string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start wisdomClient port default 9090",
	Run: func(cmd *cobra.Command, args []string) {
		if config == "" || port == 0 || log == "" {
			_ = cmd.Help()
			return
		} else {
			runStart(config, log, port)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&config, "config", "c", "./conf/config.yaml", "--config <fileLocation> example --config ./conf/config.yaml")
	runCmd.Flags().StringVarP(&log, "log", "l", "./logs/wisdom-client.log", "--log <logLocation> example --config ./logs/wisdom-client.log")
	runCmd.Flags().IntVarP(&port, "port", "p", 9090, "--port <port> example --port 9090")
}

func runStart(configFile, logPath string, port int) {
	// 初始化日志模块
	//logPath := wisdomClient.BaseDir() + "/logs/wisdom-client.log"
	err := logger.InitLogger(logPath, 1, 7, 10, "DEBUG")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 初始化配置文件
	//configPath := wisdomClient.BaseDir() + configFile
	conf.InitConfig(configFile)

	// 启动核心
	collector.StartMonitor()

	// 启动服务
	fmt.Printf("WisdomClient ListenAndServer %d...\n", port)
	r := collector.InitHttpListen()
	_ = r.Run(fmt.Sprintf(":%d", port))
}
