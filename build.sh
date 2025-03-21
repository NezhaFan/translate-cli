#!/bin/bash

# 获取系统架构和操作系统
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# 转换架构名称为 Go 使用的格式
case ${ARCH} in
    x86_64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "不支持的架构: ${ARCH}"
        exit 1
        ;;
esac

# 构建 Docker 镜像
echo "开始构建 Docker 镜像..."
echo "目标平台: linux/${ARCH}"
docker build --platform linux/${ARCH} \
  --build-arg TARGETOS=${OS} \
  --build-arg TARGETARCH=${ARCH} \
  -t tc-builder .

# 创建临时容器
echo "创建临时容器..."
docker create --name temp-container tc-builder

# 复制编译好的程序
echo "复制编译好的程序..."
docker cp temp-container:/app/tc ./tc

# 清理临时容器
echo "清理临时容器..."
docker rm temp-container

# 设置可执行权限
chmod +x ./tc

# 移动到用户的 bin 目录
echo "安装到系统..."
sudo mv ./tc /usr/local/bin/

echo "构建完成！可执行文件已安装为 'tc' 命令"
echo "现在可以直接使用 'tc curl -h' 命令了"