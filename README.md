# 今日热榜

#### 项目介绍
**鱼塘热榜是一个获取各大热门网站热门头条的聚合网站，使用Go语言编写，多协程异步快速抓取信息**

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


