## docker容器化部署MySQL数据库方法
- [ ] [部署MySQL8.0数据库]
1. 下载MySQL8.0镜像
    ```shell
    docker pull mysql:8.0.0
    ```
2. 直接运行MySQL8.0数据库
    ```shell
    docker run  -d  \
    --name mysql8-server \
    --privileged=true \
    --restart=always \
    -p 3308:3306 \
    -e MYSQL_ROOT_PASSWORD=root \
    -e TZ=Asia/Shanghai mysql:8.0.0 \
    --lower_case_table_names=1
    ```
> 参数解读说明
- `--name mysql8-server`给容器命名
- `--privileged=true` 已特权模式运行
- `--restart=always` 服务器启动时，自启动
- `-p 3308:3306` 端口映射，第一个 3308 是映射出去的端口，第二个 3306 是这个容器的端口
- `-e MYSQL_ROOT_PASSWORD=root` 环境变量设置，此处是设置 ROOT 用户登录密码
- `-e TZ=Asia/Shanghai mysql:8.0.0`  此处是设置 MySQL 的时区，请注意这点，有时候你可能会发现你的服务器时区和你当前的电脑的时区是不一样的，这很有可能有一些隐藏问题噢。此处的 mysql:8.0.0为镜像名。
- `--lower_case_table_names=1` 让表名忽略大小写
