![](img/demo1.jpg)

## [English](README.md) | 中文

## Introduction

执行命令并且使用大模型对输出进行翻译. (翻译好坏和速度主要和模型性能有关)

## Download
- 方式一: 用Docker镜像
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `docker build -t tc .` 
  - 设置环境变量 `alias tc='docker run --rm -i --tty=false -e TC_LLM_TYPE=ollama -e TC_LLM_MODEL=qwen2.5:3b -e TC_LLM_URL=http://host.docker.internal:11434 tc'` (Tip: 注意在容器内部访问本机是 `host.docker.internal` 而不是`localhost`)
- 方式二: 用Docker自己编译 (linux or mac) 
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `cd translate-cli && chmod +x build.sh && sh build.sh`
- 方式三: 你有golang环境
  - `go install github.com/NezhaFan/translate-cli@latest` 
  - `mv $GOPATH/bin/translate-cli $GOPATH/bin/tc`

## Config (方式二或方式三)
设置环境变量 (以mac举例)
```conf
# 不是ollama的话空着
export TC_LLM_TYPE=ollama
export TC_LLM_URL=http://127.0.0.1:11434
# 尽量用小的模型，别用深度思考模型。
export TC_LLM_MODEL=qwen2.5:3b
export TC_LLM_KEY=sk-
```

## Run
用管道的形式把输出传给翻译程序 `curl -h | tc`

## 可怜可怜孩子，给个🌟再走吧...
