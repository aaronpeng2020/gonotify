请编写：Dockerfile: Docker镜像构建脚本，用于将GoNotify服务打包成Docker镜像。
以下是一个简单的Dockerfile文件示例，用于将GoNotify服务打包成Docker镜像：

Dockerfile
Copy code
FROM golang:1.17 AS build
WORKDIR /app
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o GoNotify .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/GoNotify .
COPY config/config.yaml .
CMD ["./GoNotify"]
该文件定义了一个Docker镜像构建脚本，将GoNotify服务打包成Docker镜像。首先，从golang:1.17镜像中创建一个build阶段，设置工作目录为/app，并复制所有文件。在该阶段中，首先下载依赖包，然后使用CGO_ENABLED=0和GOOS=linux等参数，编译GoNotify二进制文件。

然后，从alpine:latest镜像中创建一个新的阶段。在该阶段中，首先安装ca-certificates包，以便GoNotify可以访问HTTPS资源。然后，设置工作目录为/app，并复制GoNotify二进制文件。最后，复制config.yaml文件，并设置CMD指令为./GoNotify。