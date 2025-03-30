# 1. 使用 Go 官方镜像作为构建环境
FROM golang:1.23 AS builder

# 2. 设置工作目录
WORKDIR /app

# 3. 复制 Go 项目代码
COPY . .

# 4. 下载并安装 Go 项目的依赖
RUN go mod tidy

# 5. 编译 Go 程序，生成二进制文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64  go build -o app .

# 6. 检查编译出来的二进制文件是否存在
RUN ls -la /app

# 使用更小的基础镜像运行 Go 二进制文件
FROM alpine:latest

# 复制编译好的二进制文件
COPY --from=builder /app/app /root/app

# 设置执行权限
RUN chmod +x /root/app

# 运行应用
CMD ["/root/app"]
