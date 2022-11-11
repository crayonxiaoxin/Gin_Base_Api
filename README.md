## Gin Base Api Project

#### 准备

安装命令行工具 swag，用于生成 swagger 文档
```
go get -u github.com/swaggo/swag/cmd/swag
``` 

[参考资料](https://github.com/swaggo/gin-swagger)


#### 运行
```
git clone https://github.com/crayonxiaoxin/Gin_Base_Api.git

cd Gin_Base_Api

go mod tidy

go run main.go
```


#### 生成文档
```
swag init
```

文档地址：/swagger/index.html


#### 目录结构
```
.
├── README.md
├── conf
│   └── app.conf                // 配置文件
├── controllers                 // 控制器
│   ├── login.go
│   ├── register.go
│   └── user.go
├── docs                        // 文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── main.go
├── middlewares                 // 中间件
│   └── token.go
├── models                      // 模型
│   ├── login.go
│   └── user.go
└── utils
    ├── db.go                   // 连接数据库
    ├── gorm.go                 // 自定义 GORM Model
    ├── password.go             // 密码加密解密
    ├── rc.go                   // 统一返回
    └── token.go                // jwt token
```