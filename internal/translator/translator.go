package translator

import (
	"bytes"
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
		afterTrans, err := llm.Translate(strings.Join(beforeTrans, "\n"))
		if err != nil {
			afterTrans = err.Error()
		}
		// 阻塞
		done <- printLines(strLines, beforeTrans, afterTrans)
		wg.Done()
	}()

	return done
}

// 解析
func parse(lines [][]byte) (notTrans []string, beforeTrans []string) {
	var maxLen int
	for _, line := range lines {
		if len(line) > 0 {
			s := string(line)
			idx := strings.LastIndex(s, "  ")
			if idx > -1 && len(strings.Split(s[idx+1:], " ")) <= 2 {
				idx = 0
			}
			if idx == -1 {
				idx = 0
			}
			s1 := strings.TrimSpace(s[:idx])
			notTrans = append(notTrans, s1)
			beforeTrans = append(beforeTrans, strings.TrimSpace(s[idx+1:]))
			if maxLen < len(s1) {
				maxLen = len(s1)
			}
		} else {
			notTrans = append(notTrans, "")
			beforeTrans = append(beforeTrans, "")
		}
	}

	for i, s := range notTrans {
		if maxLen < 60 && s != "" {
			notTrans[i] = s + strings.Repeat(" ", maxLen-len(s))
		}
	}
	return
}

func printLines(notTrans []string, beforeTrans []string, afterTrans string) string {
	var buff bytes.Buffer
	buff.Grow(4096)
	trans := strings.Split(afterTrans, "\n")
	for i, notTran := range notTrans {
		if len(notTran) > 0 {
			buff.WriteString(notTran)
			buff.WriteByte('\t')
		}
		if i < len(trans) && trans[i] != "" {
			buff.WriteString(trans[i])
		} else {
			buff.WriteString(beforeTrans[i])
		}
		buff.WriteByte('\n')
	}
	return buff.String()
}
