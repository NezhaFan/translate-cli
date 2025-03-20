# Translate-Cli

[English](README.md) | 中文

执行命令行并对输出进行翻译
![](demo.jpg)

# 使用方式
- 下载
  - 不会`go`的，右侧`release`下载二进制文件
  - 会`go`的 `go install github.com/NezhaFan/translate-cli@latest`
- 本地配置环境变量（举例mac）
  ``conf
  export TRANSLATE_LANG=Chinese
  export TRANSLATE_LLM_TYPE=ollama
  export TRANSLATE_LLM_URL=https://api.openai.com
  export TRANSLATE_LLM_MODEL=
  export TRANSLATE_LLM_KEY=xxxx
  ``
