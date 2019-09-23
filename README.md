# 今日热榜

[![Build Status](https://travis-ci.com/async-rs/async-std.svg?branch=master)](https://github.com/tophubs/TopList/)
[![License](https://img.shields.io/badge/license-MIT%2FApache--2.0-blue.svg)](https://github.com/tophubs/TopList/)

**今日热榜是一个获取各大热门网站热门头条的聚合网站，使用Go语言编写，多协程异步快速抓取信息，预览:[https://www.printf520.com/hot.html][热榜]**
![DeepinScrot-4337.png](https://i.loli.net/2019/08/05/PjX2nqWAgM5xsL4.png)

### 开发教程

1. 将 `/your/path/to/TopList/src` 添加到 $GOPATH，（如果用IDE开发，需要代码提示的话）
2. `cd {root_path}` # 项目根目录
3. `cp docker/dev/default.env docker/dev/.env` # 复制docker-compose环境文件
4. `make gomod` 或 `make dep` # 安装依赖包, `make gomod`是使用 `gomodule`, `make dep`是使用 `dep` 工具
5. `make dev` # 启动。会根据database.sql自动创建数据库，同时使用 [AIR](https://github.com/cosmtrek/air), 不需要重启即可热重载
6. 等待服务启动完毕，打开`http://{yourdomain}:{DEV_PORT}/` 即可访问今日热榜

### 编译教程
```
cd {root_path} # 项目根目录
make build-server
make build-gethot
```

### 目录说明

```
TopList/
├── docker
│   └── dev
│       ├── air                 # air 的可执行文件
│       ├── air-gethot.conf     # GetHot 的 air 热加载配置
│       ├── air-server.conf     # Server 的 air 热加载配置
│       ├── default.env         # docker-compose 默认的环境变量
│       └── docker-compose.yaml # docker-compose 配置文件
└── src
    └── app
        ├── App
        │   ├── GetHot.go     爬虫程序需要Cron定时任务执行
        │   └── Server.go     Server程序需要守护进程的方式执行
        ├── Common
        │   ├── Db.go         DB组件
        │   ├── Message.go
        │   └── database.sql  数据库建表文件
        ├── Config
        │   ├── MySql.go      mysql配置读取组件
        │   └── mysql.toml    mysql配置文件需要手动配置
        ├── Cron
        │   ├── GetHot.sh     爬虫Cron程序可以是每小时执行一次
        │   └── README.md
        ├── Html
        │   ├── css
        │   ├── hot.html      前端热榜展示网页
        │   └── js
        └── README.md
```

### API说明

#### 获取所有类型
- Method: **GET**
- URL:  ```https://www.printf520.com:8080/GetType```
- Param：无
- Body:
```
{
    "Code":0,
    "Message":"获取数据成功",
    "Data":[
        {"id":"1","sort":"63908","title":"知乎"},
        {"id":"2","sort":"21912","title":"虎扑"},
        {"id":"6","sort":"11707","title":"天涯"},
        {"id":"7","sort":"12546","title":"知乎日报"},
       ]}
```


### 获取具体类型热榜数据
- Method: **GET**
- URL:  ```  https://www.printf520.com:8080/GetTypeInfo?id=2```
- Param：id
- Body:
```
{
    "Code":0,
    "Message":"获取成功",
    "Data":[
        {
            "title":"
45个经典面试回答提示，分享给即将工作的大家。 zt
",
            "url":"https://bbs.hupu.com//28814429.html"
        },
        {
            "title":"
[名场面]回家的诱惑：洪世贤酒店幽会，抵不住诱惑犯了错！ zt
",
            "url":"https://bbs.hupu.com//28818367.html"
        },
        {
            "title":"
张艺兴回应假唱风波。ZT
",
            "url":"https://bbs.hupu.com//28815609.html"
        }
    ]
}
```


### 使用说明

1. fork 项目

### 参与贡献

1. Fork 本项目
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request


[热榜]: https://www.printf520.com/hot.html