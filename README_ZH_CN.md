![](img/demo1.jpg)

## [English](README.md) | 中文

## Introduction

执行命令并且使用大模型对输出进行翻译. (翻译好坏和速度主要和模型性能有关)

## Download
- 方式一: 用Docker镜像
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `docker build -t tc .` 
  - 设置环境变量 `alias tc='docker run --rm -i --tty=false -e LLM_TYPE=ollama -e LLM_MODEL=qwen2.5:3b -e LLM_URL=http://host.docker.internal:11434 tc'` (Tip: 注意在容器内部访问本机是 `host.docker.internal` 而不是`localhost`)
- 方式二: 用Docker自己编译 (linux or mac) 
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `cd translate-cli && chmod +x build.sh && sh build.sh`
- 方式三: 你有golang环境
  - `go install github.com/NezhaFan/translate-cli@latest` 
  - `mv $GOPATH/bin/translate-cli $GOPATH/bin/tc`


## Config (方式二或方式三)
- 大模型选择
  - 参数大小：本地尽量3B及以上，远程调用可以使用32B以上大一些的更出色。
  - 不要使用深度思考模型，例如 `deepseek-r1` 系列
- 推荐
  - 本地建议实测 `qwen2.5:3b` 还不错
  - 字节火山
    - 模型开通 https://console.volcengine.com/ark/region:ark+cn-beijing/openManagement?LLM=%7B%7D&OpenTokenDrawer=false&tab=LLM
    - 申请Key https://console.volcengine.com/ark/region:ark+cn-beijing/apiKey
    - 实测: `doubao-1-5-lite-32k-250115`、 `doubao-vision-lite-32k-241015`、 `doubao-1.5-pro-32k-250115` 可以
  - 阿里百炼
    - 模型列表 https://help.aliyun.com/zh/model-studio/getting-started/models
    - 申请Key https://bailian.console.aliyun.com/?apiKey=1#/api-key
    - 实测 `qwq-32b-preview` 可以
- 设置环境变量 (以mac举例)
  ```conf
  # 不是ollama的话空着
  export LLM_TYPE=ollama
  export LLM_URL=http://127.0.0.1:11434
  # 尽量用小的模型，别用深度思考模型。
  export LLM_MODEL=qwen2.5:3b
  export LLM_KEY=sk-
  ```

## Run
用管道的形式把输出传给翻译程序 `curl -h | tc`

## 可怜可怜孩子，给个🌟再走吧...
