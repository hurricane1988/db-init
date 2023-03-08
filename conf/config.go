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

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/wonderivan/logger"
)

// LoadConfigFromFile 初始化配置文件
func LoadConfigFromFile() (*Config, error) {
	// 设置viper的环境变量前缀
	viperInstance := viper.New()
	// 将viper设置为从环境变量中读取配置
	viperInstance.AutomaticEnv()
	// 从环境变量读取配置
	/* viper动态加载配置 */
	// 处理配置变更事件
	viperInstance.OnConfigChange(func(in fsnotify.Event) {
		// TODO: 联调测试
		// fmt.Println(viperInstance)
		logger.Info("配置文件" + in.Name + "发生变更")
		NewConfig(viperInstance)
	})
	NewConfig(viperInstance)
	// 监视配置文件是否发生更改
	viperInstance.WatchConfig()
	var conf = NewConfig(viperInstance)
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// NewConfig 创建config
func NewConfig(viper *viper.Viper) *Config {
	return &Config{
		MySQLConfig: *NewMySQLConfig(viper),
	}
}

// NewMySQLConfig 新的MySQL数据库配置
func NewMySQLConfig(viper *viper.Viper) *mySQLConfig {
	return &mySQLConfig{
		Host:     viper.GetString("MYSQL_HOST"),
		Username: viper.GetString("MYSQL_USERNAME"),
		Password: viper.GetString("MYSQL_PASSWORD"),
		Port:     viper.GetString("MYSQL_PORT"),
		Version:  viper.GetString("MYSQL_VERSION"),
		DB:       viper.GetString("MYSQL_DB"),
	}
}
