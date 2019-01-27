# gin_jwt_demo

这个 Dome 主要演示在 gin 框架下 jwt(json web token) 的具体实现, 同时结合 gorm 和 mysql 大体搭建了一个项目框架

## Structure

    .
    ├── controller        // 控制器
    │   └── user.go
    ├── main.go           // 入口文件                               
    ├── middleware        // 中间件
    │   └── jwt.go        // 验证 jwt
    ├── model             // 数据模型
    │   ├── database.go   // 初始化数据库
    │   └── user.go
    ├── README.md
    ├── routes            // 路由
    │   └── router.go
    ├── utils             // 工具库    
    │   └── TokenUtil.go  // 生成/解析 jwt
    └── vendor            // 依赖包 
        └── vendor.json

## How to use

### 安装依赖
    cd $GOPATH/src
    git clone git@github.com:AustinDeng/gin_jwt_demo.git
    cd gin_jwt_demo && govendor snyc

### 编辑数据库文件
    vim model/database.go

### 运行
    go run main.go

## Example

    // curl 127.0.0.1:8080/login -d username=username -d password=password
    {
      "data": {
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRlbmciLCJleHAiOjE1NDg1NzY0MDUsImlzcyI6IkF1c3RpbkRlbmcifQ.PikuZ5amw5AHV1RvA-YBeRjUN3LQf2IX0WXFtnXHguc"
      },
      "msg": "校验成功"
    }


    // curl -G  127.0.0.1:8080/api/user -d username=deng  -H 'token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRlbmciLCJleHAiOjE1NDg1NzY0MDUsImlzcyI6IkF1c3RpbkRlbmcifQ.PikuZ5amw5AHV1RvA-YBeRjUN3LQf2IX0WXFtnXHguc'
    {
      "msg": "OK",
      "user": {
          "ID": 1,
          "CreatedAt": "2019-01-23T03:59:10+08:00",
          "UpdatedAt": "2019-01-24T17:29:43+08:00",
          "DeletedAt": null,
          "Username": "deng",
          "Password": "123"
      }
    }


