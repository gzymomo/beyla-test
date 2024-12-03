# 使用官方的 Go 运行时镜像作为基础镜像
FROM golang:1.23

# 设置工作目录
WORKDIR /app

# 将当前目录的所有文件复制到容器中的工作目录
COPY . .

# # 下载并安装应用所需的第三方依赖
# RUN go mod download
# RUN go mod verdor

# 编译 Go 程序
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -o beylaapp  .

# 使用最小的基础镜像运行应用
FROM alpine:3.15

# 设置工作目录
WORKDIR /opt/

# 从构建阶段复制 Go 二进制文件
COPY --from=0 /app/beylaapp .
RUN chmod +x /opt/beylaapp

# 设置容器启动时运行的命令
CMD ["/opt/beylaapp"]
