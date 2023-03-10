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
	"strings"
)

// DbInitHandler 数据库初始化处理方法
func DbInitHandler() {
	// 加载全局config
	var cfg = conf.GlobalConfig
	// 如果数据库类型是MySQL数据库
	if strings.EqualFold(cfg.Type, conf.MySQL) {

	}
}
