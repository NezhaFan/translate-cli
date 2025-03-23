package translator

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/NezhaFan/translate-cli/internal/model"
)

func TransAndPrint(lines [][]byte, llm model.LLM, wg *sync.WaitGroup) chan string {
	wg.Add(1)
	// 提取需要翻译片段
	strLines, beforeTrans := parse(lines)

	// 并发去翻译
	done := make(chan string)
	go func() {
		afterTrans, err := llm.Translate(beforeTrans)
		if err != nil {
			fmt.Printf("翻译出错: %v\n", err)
			os.Exit(1)
		}
		// 阻塞
		done <- printLines(strLines, afterTrans)
		wg.Done()
	}()

	return done
}

var helpRegex = regexp.MustCompile(`(?m)^(\s*(?:(?:-\w+|--[\w-]+)(?:,\s*(?:-\w+|--[\w-]+))*\s*(?:<[^>]+>)?\s*)?)(.+)$`)

// 解析
func parse(lines [][]byte) (rows []string, beforeTrans string) {
	var buff bytes.Buffer
	buff.Grow(4096)
	var maxLen int
	for _, line := range lines {
		matches := helpRegex.FindSubmatch(line)
		if len(matches) > 2 {
			buff.Write(matches[2])
		}
		if maxLen < len(line) {
			maxLen = len(line)
		}
		buff.WriteByte('\n')

	}

	beforeTrans = buff.String()
	for _, line := range lines {
		rows = append(rows, string(line)+strings.Repeat(" ", maxLen-len(line)))
	}
	return
}

func printLines(rows []string, afterTrans string) string {
	var buff bytes.Buffer
	buff.Grow(4096)
	trans := strings.Split(afterTrans, "\n")
	for i, line := range rows {
		buff.WriteString(line)
		buff.WriteByte('\t')
		if i < len(trans) {
			buff.WriteString(trans[i])
		}
		buff.WriteByte('\n')
	}
	return buff.String()
}
