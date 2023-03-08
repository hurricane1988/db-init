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

package mysql

import (
	"sync"
)

// 定义MySQL版本常量
const (
	MysqlV5 = "5"
	MysqlV8 = "8"
	// SetConnMaxLifetime 设置数据库最大连接数
	SetConnMaxLifetime = 100
	// SetMaxIdleConns 设置上数据库最大闲置连接数
	SetMaxIdleConns = 10
)

// MySQL 定义数据库链接结构体
type MySQL struct {
	Host     string     `json:"host" description:"数据库链接地址"`
	Username string     `json:"username" description:"数据库链接账号"`
	Password string     `json:"password" description:"数据库链接密码"`
	Port     string     `json:"port" description:"数据库链接端口"`
	Version  string     `json:"version" description:"MySQL数据库版本"`
	Mutex    sync.Mutex `json:"mutex" description:"MySQL数据库互斥锁"`
	DB       string     `json:"db" description:"MySQL数据库DB"`
}
