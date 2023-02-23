# GoNotify - 平台余额通知服务

## 项目需求
服务可以定时请求多个平台的接口来查询平台的余额，如果余额低于配置的预警值，则发送邮件和短信进行提醒，接下来我会提供每个平台的接口文档供开发，我希望服务通过docker来部署，同时可以不断添加新的平台进来。

## 项目架构
服务端：
使用 Go语言 编写服务端程序
使用 Cron 或者类似的调度工具来定期请求（10分钟）平台余额信息
无需操作数据库
集成发送腾讯云的邮件和短信SDK
根据余额和预警值来发送提醒
使用 Docker 打包和编排项目文件
提供配置文件，用于设置平台接口信息、预警值、发送邮件和短信的配置等

平台接口：
每个平台的接口可能有不同的请求参数和返回格式，需要根据接口文档进行开发和集成

