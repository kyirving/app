# 1、基础镜像 Go 1.26 Alpine 镜像， 别名 builder 镜像
FROM golang:1.26.2-alpine as builder

# RUN 构建过程执行
# CMD 运行时执行

# 设置构建环境变量
ENV GO111MODULE=on \
   GOOS=linux \
   GOARCH=amd64 \
   CGO_ENABLED=0

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum（利用 Docker 缓存，依赖不常变）
COPY go.mod go.sum ./
RUN go mod download

# 复制项目代码到工作目录
COPY . .

# 编译项目 可替换编译入口和输出文件名
RUN go build -ldflags="-s -w" -o ./bin/server/myapp ./cmd/server/main.go


# 2、运行镜像
FROM alpine:3.22.4 as runner

# 安装 ca-certificates 和时区数据（方便处理 HTTPS 请求和时区）
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/bin/server/myapp .
COPY --from=builder /app/config .

# 暴露应用的端口（根据实际应用修改）
EXPOSE 8080

# 启动应用
CMD ["./myapp", "--configDir", "./config"]
