# 使用官方的Go镜像作为基础镜像
FROM golang:1.20-alpine

# 设置Go模块代理为阿里云源
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将源代码复制到工作目录
COPY . .

# 构建应用程序
RUN go build -o myapp ./cmd/myapp

# 暴露应用程序运行的端口
EXPOSE 8080

# 运行应用程序
CMD ["./myapp"]
