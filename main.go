package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"tc/model"
)

var (
	config struct {
		Type string
		model.Config
	}
)

func init() {
	config.Lang = os.Getenv("TC_LANG")
	config.Url = os.Getenv("TC_LLM_URL")
	config.Model = os.Getenv("TC_LLM_MODEL")
	config.Key = os.Getenv("TC_LLM_KEY")
	config.Type = os.Getenv("TC_LLM_TYPE")
	if config.Lang == "" {
		config.Lang = "Chinese"
	}
}

func main() {
	if config.Url == "" || config.Model == "" || (config.Key == "" && config.Type != "ollama") {
		fmt.Println("Please set environment variables according to the doc.")
		return
	}

	// nothing
	commands := os.Args[1:]
	if len(commands) == 0 {
		fmt.Println("Just translate!")
		return
	}

	// 执行命令并获取输出
	output, err := executeCommand(commands...)
	if err != nil {
		fmt.Printf("命令执行失败: %v\n", err)
		return
	}
	if len(output) > 1024*8 {
		fmt.Println(output)
		fmt.Println("输出过长，不进行翻译")
		return
	}

	// 分离出需要翻译部分
	lines, beforeTrans := parse(output)
	// fmt.Println(beforeTrans)

	// 模型翻译
	llm := model.New(config.Type, config.Config)
	afterTrans, err := llm.Translate(beforeTrans)
	if err != nil {
		fmt.Printf("翻译出错: %v\n", err)
		return
	}

	trans := strings.Split(afterTrans, "\n")
	for i := range lines {
		if i < len(trans) {
			fmt.Printf("%s \t %s \n", lines[i], trans[i])
		} else {
			fmt.Printf("%s \t\n", lines[i])
		}
	}
}

// 执行命令并获取输出
func executeCommand(cmds ...string) (string, error) {
	// 检查命令是否存在
	path, err := exec.LookPath(cmds[0])
	if err != nil {
		return "", fmt.Errorf("命令 '%s' 未找到: %v", strings.Join(cmds, " "), err)
	}
	cmd := exec.Command(path, cmds[1:]...)
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		return string(output), nil
	}
	return "", fmt.Errorf("%s: %s %v", strings.Join(cmds, " "), output, err)
}

var helpRegex = regexp.MustCompile(`(?m)^(\s*(?:(?:-\w+|--[\w-]+)(?:,\s*(?:-\w+|--[\w-]+))*\s*(?:<[^>]+>)?\s*)?)(.+)$`)

func parse(output string) (lines []string, trans string) {
	buf := bytes.NewBuffer(nil)
	buf.Grow(len(output))

	var maxLen int
	lines = strings.Split(output, "\n")
	for _, line := range lines {
		matches := helpRegex.FindStringSubmatch(strings.TrimSpace(line))
		if len(matches) > 2 {
			buf.WriteString(strings.TrimSpace(matches[2]))
			if maxLen < len(line) {
				maxLen = len(line)
			}
		}
		buf.WriteRune('\n')
	}
	trans = buf.String()

	// 对齐
	for i := range lines {
		lines[i] += strings.Repeat(" ", maxLen-len(lines[i]))
	}
	return
}
