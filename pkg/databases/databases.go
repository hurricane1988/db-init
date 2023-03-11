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

package databases

import (
	"db-init/conf"
	"db-init/utils/db/mysql"
	"errors"
	"github.com/wonderivan/logger"
	"strings"
	"sync"
)

// DBInitHandler 数据库初始化处理方法
func DBInitHandler(filePath string) error {
	// 加载全局config
	var (
		options = conf.GlobalConfig.MySQLConfig
	)
	// 如果数据库类型是MySQL数据库
	if strings.EqualFold(conf.GlobalConfig.Type, conf.MySQL) {
		mysqlClient, err := mysql.NewClient(
			options.Host,
			options.Username,
			options.Password,
			options.Port,
			options.Database,
			options.Version,
			sync.Mutex{})
		if err != nil {
			panic(err)
			return nil
		}
		/* 初始化MySQL数据库 */
		// 初始化数据库失败，执行SQL异常
		if err = mysqlClient.Exec(filePath); err != nil {
			logger.Error("执行SQL文件 "+filePath+"失败,错误信息", err)
			return err
		}
		// 初始化数据库成功
		return nil
	} else {
		return errors.New("数据库版本不执行")
	}
}
