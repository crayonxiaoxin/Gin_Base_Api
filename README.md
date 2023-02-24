## Gin Base Api Project

### 下载
```
git clone https://github.com/crayonxiaoxin/Gin_Base_Api.git
```


### 准备

配置数据库 conf.ini（值不需要引号）
```
; 监听端口
port = 8083

; 开发环境
[dev]
db_user = 
db_pass = 
db_host = 
db_name = 
```

安装命令行工具 swag，用于生成 swagger 文档 ([参考资料](https://github.com/swaggo/gin-swagger))
```
go get -u github.com/swaggo/swag/cmd/swag
``` 


### 获取依赖
```
go mod tidy
```

### 运行
```
dev=1 go run .
```


### 生成&更新文档
```
swag init
```

文档地址：/swagger/index.html 或 /docs



### 生产环境 build
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build . 
```


### 部署
> 将 可执行文件 以及 conf.ini 上传到服务器

> 通过 Apache 或 Nginx 反代即可配置 domain

Nginx 配置示例：
```
server
{
    listen 80;
    server_name api.example.xyz;
    index index.html index.htm default.htm default.html;
    root /var/www/gin/hello_gin_api;

    # HTTP反向代理相关配置开始 >>>
    location ~ /purge(/.*) {
        proxy_cache_purge cache_one 127.0.0.1$request_uri$is_args$args;
    }

    location / {
        proxy_pass http://127.0.0.1:8083;
        proxy_set_header Host 127.0.0.1:$server_port;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header REMOTE-HOST $remote_addr;
        add_header X-Cache $upstream_cache_status;
        proxy_set_header X-Host $host:$server_port;
        proxy_set_header X-Scheme $scheme;
        proxy_connect_timeout 30s;
        proxy_read_timeout 86400s;
        proxy_send_timeout 30s;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
    # HTTP反向代理相关配置结束 <<<
}
```


### 目录结构
```
.
├── README.md
├── conf.ini                    // 配置文件
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