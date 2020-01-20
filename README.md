# 今日热榜

**今日热榜是一个获取各大热门网站热门头条的聚合网站，使用Go语言编写，多协程异步快速抓取信息，预览:[http://mo.fish][热榜]**

![DeepinScrot-1130.png](http://img.printf520.com/img/DeepinScrot-1130.png)

### 安装教程

1. 编译

   ```
   cd {root_path} # 项目根目录
   go build -o ./App/GetHot App/GetHot.go
   go build -o ./App/Server App/Server.go 
   ```
   
2. 创建数据库，如 `news`，执行database.sql创建表，更改配置文件`Config/mysql.toml`

3. 编辑文件 `Html/js/blog/globalConfig.js`

   ```
   const ServerIp = 'http://{your_domain}:9090' // 替换成服务器域名
   ```

4. 部署定时任务/App/GetHot.go爬虫程序，且以守护进程的方式执行Server.go

   ```
   crontab -e # 添加一行 0 */1 * * * {root_path}/App/GetHot
   nohup {root_path}/App/Server &
   ```

5. 测试

   - 打开`http://{yourdomain}:9090/` 即可访问今日热榜


### 目录说明

```
TopList/
├── App
│   ├── GetHot.go   爬虫程序需要Cron定时任务执行
│   └── Server.go   Server程序需要守护进程的方式执行
├── Common
│   ├── Db.go       DB组件
│   └── Message.go  
├── Config
│   ├── MySql.go    mysql配置读取组件
│   └── mysql.toml  mysql配置文件需要手动配置
├── Cron
│   ├── GetHot.sh   爬虫Cron程序可以是每小时执行一次
│   └── README.md
├── database.sql    数据库建表文件
├── Html
│   ├── css
│   ├── hot.html    前端热榜展示网页
│   └── js
│  
└── README.md
```

### API说明

#### 获取所有类型
- Method: **GET**
- URL:  ```https://www.tophub.fun:8888/GetAllType```
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
- URL:  ```  https://www.tophub.fun:8888/v2/GetAllInfoGzip?id=59&page=0```
- Param：id
- Body:
```
{
    "Code":0,
    "Message":"获取数据成功",
    "Data":{
        "data":[
            {
                "id":1327371,
                "CreateTime":1579487422,
                "commentNum":0,
                "approvalNum":0,
                "Title":"180W 买了个小破房，月供 7500 多，压力山大",
                "hotDesc":"200条评论",
                "Url":"https://www.v2ex.com/t/639366#reply200",
                "imgUrl":"",
                "isRss":"",
                "is_agree":0,
                "TypeName":"社区"
            },
            {
                "id":1326304,
                "CreateTime":1579483806,
                "commentNum":0,
                "approvalNum":0,
                "Title":"口罩(N95) 目前应该买哪个? 有研究过得推荐一下?",
                "hotDesc":"106条评论",
                "Url":"https://www.v2ex.com/t/639343#reply106",
                "imgUrl":"",
                "isRss":"",
                "is_agree":0,
                "TypeName":"社区"
            },
        ],
        "page":10
    }
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
