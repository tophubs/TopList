# 鱼塘热榜

**鱼塘热榜是一个获取各大热门网站热门头条的聚合网站，使用Go语言编写，多协程异步快速抓取信息，预览:https://www.printf520.com/hot.html**
![DeepinScrot-4337.png](https://i.loli.net/2019/08/05/PjX2nqWAgM5xsL4.png)

#### 安装教程

1. git clone https://github.com/timeromantic/hotDataList.git
2. 执行database.sql文件
3. 配置/Config/Mysql.go数据库地址
4. 执行/App/GetHot.go程序
5. 修改hot.html并打开
6. 部署定时任务


#### 使用说明

1. fork 项目

#### 参与贡献

1. Fork 本项目
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request

#### 目录说明

```
HotList/
├── App
│   └── GetHot.go  爬虫主程序
├── Common
│   ├── Db.go      数据库组件
│   └── Redis.go   redis组件
├── Config
│   ├── Config.go 
│   └── Mysql.go   mysql配置文件
├── Cron
│   └── GetHot.sh  爬虫定时脚本
├── Exe
├── Html
│   ├── css
│   ├── hot.html   热榜展示网页
│   └── js
|
└── database.sql
└── README.md
```

#### API说明

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


#### 获取具体类型热榜数据
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



