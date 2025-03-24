![](img/demo1.jpg)

##  English | [中文](README.md)


## Introduction

Execute commands and use LLM to translate output.


## Download
- Option 1: Only use **docker**
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `docker build -t tc .` 
  - Set to short environment variables. for example: `alias tc='docker run --rm -i --tty=false -e LLM_TYPE=ollama -e LLM_MODEL=qwen2.5:3b -e LLM_URL=http://host.docker.internal:11434 tc'` (Tip: When you access a native address inside a container, you should use `host.docker.internal` instead of `localhost`)
- Option 2: Compile it yourself with **docker** (linux or mac) 
  - `git clone https://github.com/NezhaFan/translate-cli.git`
  - `cd translate-cli && chmod +x build.sh && sh build.sh`
- Option 3: if you r a gopher
  - `go install github.com/NezhaFan/translate-cli@latest` 
  - `mv $GOPATH/bin/translate-cli $GOPATH/bin/tc`

## Config (When Option 2 or Option 3)
set env (for example in mac)
```conf
# tradnslated language
export LANG=Chinese
# ollama or empty
export LLM_TYPE=ollama
export LLM_URL=http://127.0.0.1:11434
# use small model, don't use deep thinking model
export LLM_MODEL=qwen2.5:3b
# if not ollama it is required
export LLM_KEY=sk-
```

## Run
Transmission by pipeline `curl -h | tc`

## Give me star 🌟！ Please!
