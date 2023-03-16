#Copyright 2022 qkp Authors
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.


.PHONY: build-binary remove-binary gotool build-image save-image remove-image

all: build-binary remove-binary gotool build-image save-image remove-image

build-binary: ;$(info $(M)...Begin to build binary files.)  @ ## Build the binary file
	hack/lib/build_binary.sh

remove-binary: ;$(info $(M)...Begin to clean binary files.)  @ ## Clean the binary file
	hack/lib/clean_binary.sh

gotool:
	go fmt ./
	go vet ./

build-image: ;$(info $(M)...Begin to build the docker image.)  @ ## Build the docker image
	hack/lib/build_image.sh

save-image: ;$(info $(M)...Begin to save the docker image.)  @ ## Save the docker image
	hack/lib/save_image.sh

remove-image: ;$(info $(M)...Begin to remove the docker image.)  @ ## Remove the docker image
	hack/lib/remove_image.sh
help:
	@echo "-----------------------------------------------------------------------------------"
	@echo "#    make all                              - 执行所有make操作"
	@echo "#    make build-binary                     - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "#    make remove-binary                    - 移除二进制文件和 vim swap files"
	@echo "#    make gotool                           - 运行 Go 工具 'fmt' and 'vet'"
	@echo "#    make build-image                      - 构建image镜像"
	@echo "#    make save-image                       - 生成镜像tar文件"
	@echo "#    make remove-image                     - 删除image压缩文件"
	@echo "-----------------------------------------------------------------------------------"