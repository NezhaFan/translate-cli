
# Translate-Cli

English | [中文](README_ZH_CN.md)

Execute commands and use LLM to translate output.


### Download
- download program from `relase` , rename as `tc`, set as global envirenment.
- Or if you are a gopher, just `go install github.com/NezhaFan/translate-cli@latest`

### Config
set env (for example in mac)
```conf
export TC_LANG=Chinese #翻译后的语言
export TC_LLM_TYPE=ollama  #非ollama可以为空
export TC_LLM_URL=http://127.0.0.1:11434
export TC_LLM_MODEL=qwen2.5:3b #使用较小的模型速度快
export TC_LLM_KEY= #非ollama需要
```

# Use
![](img/demo1.jpg)