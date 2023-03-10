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
	"db-init/utils/db/mysql/versions"
	v5 "db-init/utils/db/mysql/versions/v5"
	"fmt"
	"github.com/wonderivan/logger"
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

// Client 定义client客户端结构体
type Client struct {
	host     string          `json:"host" xml:"host" yaml:"host"`
	username string          `json:"username,omitempty" xml:"username" yaml:"username"`
	password string          `json:"password" xml:"password" yaml:"password"`
	port     string          `json:"port" xml:"port" yaml:"port"`
	database string          `json:"database" xml:"database" yaml:"database"`
	version  string          `json:"version" xml:"version" yaml:"version"`
	client   versions.Client `json:"client" yaml:"c" xml:"c"`
	mux      sync.Mutex      `json:"mux" xml:"mux" yaml:"mux"`
}

// NewClient 初始化数据库客户端
func NewClient(host, username, password, port, database, version string, mutex sync.Mutex) (*Client, error) {
	var err error
	m := &Client{
		host:     host,
		username: username,
		password: password,
		version:  version,
		database: database,
		port:     port,
	}
	switch m.version {
	case MysqlV5:
		config := v5.MySQL{
			Host:     host,
			Username: username,
			Password: password,
			Version:  version,
			Port:     port,
			Database: database,
			Mutex:    mutex,
		}
		m.client, err = v5.New(&config)
	default:
		return nil, fmt.Errorf("不支持的数据库版本 %s", m.version)
	}
	return m, err
}

// LoadClient 加载客户端
func (c *Client) LoadClient() error {
	// 检查数据库客户端是否已经初始化
	if c.client != nil {
		return nil
	}
	// 创建数据库客户端
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.client != nil {
		return nil
	}
	var (
		vc  versions.Client
		err error
	)

	// 判断不同MySQL版本
	switch c.version {
	case MysqlV5:
		vc, err = v5.New(&v5.MySQL{Host: c.host, Username: c.username, Password: c.password, Port: c.port, Version: c.version, Database: c.database, Mutex: c.mux})
	default:
		err = fmt.Errorf("unsupported MySQL version %s", c.version)
	}
	if err != nil {
		return err
	}
	c.client = vc
	return nil
}

// Exec 执行SQL操作
func (c *Client) Exec(sqlFile string) error {
	var err error
	err = c.LoadClient()
	if err != nil {
		return err
	}
	err = c.client.Exec(sqlFile)
	if err != nil {
		logger.Error("执行sql文件失败,错误信息", err)
	}
	logger.Info("初始化数据库成功!")
	return nil
}
