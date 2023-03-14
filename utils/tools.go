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
	"db-init/conf"
	"fmt"
	"github.com/wonderivan/logger"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
)

// InitMysqlConfig 初始化MySQL配置
func InitMysqlConfig(config conf.Options) string {
	// 格式化数据库连接并返回
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
}

// LoadSuffixFile 读取sql文件
func LoadSuffixFile(directory, suffix string) (string, error) {
	// 定义局部变量
	var (
		files     []string
		fileNames []string
	)
	// 遍历指定目录下所有的文件
	err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		logger.Error(err)
	}
	// 判断文件是否包含指定的文件后缀
	for _, file := range files {
		if !reflect.DeepEqual(path.Ext(file), suffix) {
			continue
		}
		fileNames = append(fileNames, file)
	}
	// 文件内容读取
	content, err := ReadFiles(fileNames)
	if err != nil {
		logger.Error(err)
	}
	return content, nil
}

// ReadFiles 读取文件内容并返回string
func ReadFiles(files []string) (string, error) {
	// 定义全局变量
	var (
		contentBytes []byte
	)
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		// []byte append至[]byte
		contentBytes = append(contentBytes, content...)
		// 读取文件添加换行符
		contentBytes = append(contentBytes, []byte("\n")...)
	}
	// byte转为string
	return string(contentBytes), nil
}
