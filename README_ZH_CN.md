
[English](README.md) | 中文

## Translate-Cli

执行命令并且使用大模型对输出进行翻译


## Download
- 从 `relase` 下载二进制程序, 重命名为 `tc`, 设置为全局环境变量
- 或者如果你有`golang`环境 `go install github.com/NezhaFan/translate-cli@latest` 然后 `mv $GOPATH/bin/translate-cli $GOPATH/bin/tc`

## Config
设置环境变量 (以mac举例)
```conf
export TC_LANG=Chinese #翻译后的语言
export TC_LLM_TYPE=ollama  #非ollama可以为空
export TC_LLM_URL=http://127.0.0.1:11434
export TC_LLM_MODEL=qwen2.5:3b #使用较小的模型速度快
export TC_LLM_KEY=
```

## Use
![](img/demo1.jpg)

## Give me star 🌟！ Please!