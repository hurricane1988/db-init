/*
Copyright 2023 QKP Authors

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
	"db-init/conf"
	"db-init/pkg/databases"
	"db-init/utils/terminal"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
)

const (
	// Parameter 第一个参数默认值
	Parameter = "start"
	// 默认读取SQL文件路径
	FilePath = "./"
)

// InitDB 执行数据库初始化
var InitDB = &cobra.Command{
	Use:              "start",
	Version:          conf.Version,
	Example:          `db-init start`,
	Short:            "start db-init client.",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		// 终端颜色打印
		terminal.TerminalColor()
		// 判断是否为start开始执行
		if len(args[0]) == 0 {
			logger.Warn("为设置start开始参数")
		} else {
			switch args[0] {
			case Parameter:
				databases.DBInitHandler()
			default:
				logger.Warn("输入的参数有误")
			}
		}
	},
}

// 将方法注册到rootCmd
func init() {
	rootCmd.AddCommand(InitDB)
}
