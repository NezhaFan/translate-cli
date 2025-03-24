# 使用官方 Go 镜像作为构建阶段
FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制源代码
COPY . .

# 构建应用
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o tc

# 设置入口点
ENTRYPOINT ["/app/tc"]