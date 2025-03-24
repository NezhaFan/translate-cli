![](img/demo1.jpg)

## [English](README_EN.md) | 中文

## 介绍

执行命令并且使用大模型对输出进行翻译. (翻译好坏和速度主要和模型性能有关)

## 下载
- 方式一: 用Docker镜像
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `docker build -t tc .` 
  - 设置环境变量 `alias tc='docker run --rm -i --tty=false LLM_URL=https://ark.cn-beijing.volces.com/api/v3 -e LLM_KEY=xxx -e LLM_MODEL=doubao-1.5-pro-32k-250115 tc'`
- 方式二: 用Docker自己编译 (linux or mac) 
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `cd translate-cli && chmod +x build.sh && sh build.sh`
- 方式三: 你有golang环境
  - `go install github.com/NezhaFan/translate-cli@latest` 
  - `mv $GOPATH/bin/translate-cli $GOPATH/bin/tc`

## 配置 
- 大模型选择
  - 尽量使用豆包、千问这种中文模型
  - 参数大小可以多尝试几个，我测试32B大多表现OK
  - 不要使用深度思考模型，例如 `deepseek-r1` 系列
- 推荐
  - 字节火山
    - 模型开通 https://console.volcengine.com/ark/region:ark+cn-beijing/openManagement?LLM=%7B%7D&OpenTokenDrawer=false&tab=LLM
    - 申请Key https://console.volcengine.com/ark/region:ark+cn-beijing/apiKey
    - 实测: `doubao-1-5-lite-32k-250115`、 `doubao-vision-lite-32k-241015`、 `doubao-1.5-pro-32k-250115` 可以
  - 阿里百炼
    - 模型列表 https://help.aliyun.com/zh/model-studio/getting-started/models
    - 申请Key https://bailian.console.aliyun.com/?apiKey=1#/api-key
    - 实测 `qwq-32b-preview` 可以
- 设置环境变量 (方式二或方式三) (以mac举例)
  ```conf
  export LLM_URL=https://ark.cn-beijing.volces.com/api/v3
  export LLM_MODEL=doubao-1.5-pro-32k-250115
  export LLM_KEY=xxx
  ```
- 其实也支持本地ollama，如果你会的话加个 `LLM_TYPE=ollama` 就可以，但是实测有些不稳定，难以兼容准确性和速度。

## 使用
用管道的形式把输出传给翻译程序 `curl -h | tc`

## 可怜可怜孩子，给个🌟再走吧...
