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

package utils

import (
	ms "db-init/utils/db/mysql"
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
)

// InitMysqlConfig 初始化MySQL配置
func InitMysqlConfig(config *ms.MySQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DB)
}

// ReadSqlFile 读取sql文件
func ReadSqlFile(sql string) (string, error) {
	sqlByte, err := ioutil.ReadFile(sql)
	if err != nil {
		logger.Error("读取SQL文件失败,")
		return "", err
	}
	return string(sqlByte), nil
}
