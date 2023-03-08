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

package v5

import (
	"database/sql"
	"db-init/utils"
	ms "db-init/utils/db/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wonderivan/logger"
	"sync"
)

// DriverName 定义数据库driver名称
const (
	DriverName = "mysql"
)

// MysqlDB 数据库客户端结构体
type MysqlDB struct {
	DB    *sql.DB    `json:"db" description:"MySQL数据库输出化"`
	Mutex sync.Mutex `json:"mutex" description:"MySQL数据库互斥锁"`
}

// New 初始化MySQL数据库
func New(config *ms.MySQL) (*MysqlDB, error) {
	var (
		db  *sql.DB
		err error
	)
	// 连接MySQL数据库
	db, _ = sql.Open(DriverName, utils.InitMysqlConfig(config))
	db.SetConnMaxLifetime(ms.SetConnMaxLifetime)
	db.SetMaxIdleConns(ms.SetMaxIdleConns)
	// 验证数据库链接
	if err = db.Ping(); err != nil {
		logger.Error("打开数据库失败,错误信息 ", err.Error())
		return nil, err
	}
	// 关闭MySQL连接
	defer db.Close()
	return &MysqlDB{db, sync.Mutex{}}, nil
}

// Exec MySQL写入操作
func (m *MysqlDB) Exec(sqlFile string) {
	_, err := m.DB.Exec(sqlFile)
	if err != nil {
		logger.Error("执行sql文件失败,错误信息", err)
	}
	logger.Info("初始化数据库成功!")
}
