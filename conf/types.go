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

package conf

// Config 定义MySQL数据库全局配置结构体
type Config struct {
	Type        string      `json:"type" yaml:"type" xml:"type" description:"数据库类型" example:"mysql\oracle\"`
	MySQLConfig mySQLConfig `json:"mySQLConfig" yaml:"mySQLConfig" xml:"mySQLConfig"`
}

type mySQLConfig struct {
	Host     string `json:"host" description:"数据库链接地址" example:"10.10.10.110"`
	Username string `json:"username" description:"数据库链接账号" example:"root"`
	Password string `json:"password" description:"数据库链接密码" example:"password"`
	Port     string `json:"port" description:"数据库链接端口" example:"3306"`
	Version  string `json:"version,omitempty" description:"MySQL数据库版本" example:"v5"`
	Database string `json:"database,omitempty" description:"MySQL数据库DB" example:"mysql"`
}

// Options 定义数据库链接结构体
type Options struct {
	Host     string `json:"host" xml:"host" yaml:"host"description:"数据库链接地址"`
	Username string `json:"username" xml:"username" yaml:"username"description:"数据库链接账号"`
	Password string `json:"password" xml:"password" yaml:"password" description:"数据库链接密码"`
	Port     string `json:"port" xml:"port" yaml:"port"description:"数据库链接端口"`
	Version  string `json:"version" xml:"version" yaml:"version"description:"MySQL数据库版本"`
	// Mutex    sync.Mutex `json:"mutex,omitempty" xml:"mutex" yaml:"mutex" description:"MySQL数据库互斥锁"`
	Database string `json:"database,omitempty" xml:"database" yaml:"database"description:"MySQL数据库DB"`
}
