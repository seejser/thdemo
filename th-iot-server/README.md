# th-iot-server

用 go+gin+redis +邮箱验证码+图片验证码+mysql+jwt+Onenet实现物联网系统

## 创建项目

```sh
#创建项目th-iot-server
go mod init th-iot-server

#安装检查
go mod tidy

# 运行
go run main.go
```

## 测试

```sh
# 发送验证码接口
curl -X POST http://localhost:9090/send -H "Content-Type: application/json" -d '{"email": "seejser@gmail.com"}'
```
