# GoNotify - 平台余额通知服务
## 项目需求
服务可以定时请求多个平台的接口来查询平台的余额，如果余额低于配置的预警值，则发送邮件和短信进行提醒。接下来我会提供每个平台的接口文档供开发，我希望服务通过docker来部署，同时可以不断添加新的平台进来。

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
## 项目文件
项目文件结构如下：
```
go
Copy code
.
├── Dockerfile
├── README.md
├── config
│   └── config.yaml
├── go.mod
├── go.sum
├── main.go
├── platform
│   ├── platform.go
│   └── platform_test.go
├── schedule
│   └── schedule.go
├── service
│   ├── service.go
│   └── service_test.go
└── util
    ├── email.go
    ├── email_test.go
    ├── sms.go
    └── sms_test.go
```
- Dockerfile：Docker镜像构建脚本，用于将GoNotify服务打包成Docker镜像
- README.md：项目文档，包含项目需求、架构和文件结构等信息
- config/config.yaml：配置文件，包含平台接口信息、预警值、发送邮件和短信的配置等，用于配置GoNotify服务的行为
- main.go：服务入口文件，用于启动GoNotify服务
- platform/platform.go：平台接口封装文件，用于请求平台接口并解析返回结果
- chedule/schedule.go：定时任务调度文件，用于定期请求平台余额信息
- service/service.go：服务逻辑文件，包含服务的核心逻辑，如发送邮件和短信进行提醒等
- util/email.go：发送邮件工具文件，封装了发送邮件的逻辑
- util/sms.go：发送短信工具文件，封装了发送短信的逻辑
## 如何使用
1. 克隆本项目：git clone https://github.com/yourname/gonotify.git
2. 修改配置文件：`cd gonotify && vim config/config.yaml
3. 构建Docker镜像：docker build -t gonotify .
4. 启动Docker容器：docker run -d --name gonotify -v /path/to/config.yaml:/app/config.yaml gonotify
5. 查看容器日志：docker logs -f gonotify

在第3步中，使用Dockerfile将GoNotify服务打包成Docker镜像。

在第4步中，使用docker run命令启动Docker容器，并将config.yaml文件挂载到容器内部的/app目录下，以便服务程序可以读取该配置文件。

在第5步中，使用docker logs命令查看容器的日志输出，以便了解服务的运行情况。

用户可以按照以上步骤，快速使用GoNotify服务。