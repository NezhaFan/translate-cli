
# Translate-Cli

English | [中文](README_ZH_CN.md)

Execute commands and use LLM to translate output.


### Download
- download program from `relase` , rename as `tc`, set as global envirenment.
- Or if you are a gopher, just `go install github.com/NezhaFan/translate-cli@latest` and `mv $GOPATH/bin/translate-cli $GOPATH/bin/tc`

### Config
set env (for example in mac)
```conf
export TC_LANG=Chinese #Translated language
export TC_LLM_TYPE=ollama  #ollama or empty
export TC_LLM_URL=http://127.0.0.1:11434
export TC_LLM_MODEL=qwen2.5:3b #Use small model as possible
export TC_LLM_KEY=
```

# Use
![](img/demo1.jpg)