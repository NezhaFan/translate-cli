![](img/demo1.jpg)

English | [中文](README_ZH_CN.md)


## Introduction

Execute commands and use LLM to translate output.


## Download
- Option 1: Compile it yourself with **docker**
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `cd translate-cli && chmod +x && sh build.sh`
- Option 2: if you r a gopher
  - `go install github.com/NezhaFan/translate-cli@latest` 
  - `mv $GOPATH/bin/translate-cli $GOPATH/bin/tc`

## Config
set env (for example in mac)
```conf
export TC_LANG=Chinese #Translated language
export TC_LLM_TYPE=ollama  #ollama or empty
export TC_LLM_URL=http://127.0.0.1:11434
export TC_LLM_MODEL=qwen2.5:3b #Use small model as possible
export TC_LLM_KEY=sk-
```

## Give me star 🌟！ Please!
