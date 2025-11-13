# th-iot-server

用 go+gin+redis +邮箱验证码+图片验证码+mysql+jwt+Onenet 实现物联网系统

## 创建项目

```sh
#创建项目th-iot-server
go mod init th-iot-server

#安装检查
go mod tidy

#先配置环境，请按照环境配置步骤操作

# 运行
go run main.go

# 打包部署
go build ./
```

## 环境配置步骤

```sh
# 在项目目录下创建 .env文件
touch .env
#写入如下内容

# Gin 配置 release/debug
GIN_MODE=debug
# Redis
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# MySQL
MYSQL_ADDR=127.0.0.1:3306
MYSQL_USER=root
MYSQL_PASSWORD=aA123456
MYSQL_DB=thiotdb

# SMTP
SMTP_HOST=smtp.qq.com
SMTP_PORT=587
SMTP_USER=159825@qq.com
SMTP_PASS=nfuzpssr

# OneNET
ONENET_PRODUCT_ID=Ay3w00GD25
ONENET_PRODUCT_ACCESS_KEY=w7G5OVd5u9/BD+l/42FtbYcJe
ONENET_VERSION=2022-05-01
ONENET_METHOD=sha1

```

## 测试

```sh
# 发送验证码接口
curl -X POST http://localhost:9090/send -H "Content-Type: application/json" -d '{"email": "see@gmail.com"}'
```
