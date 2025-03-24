package translator

import (
	"bytes"
	"sync"

	"github.com/NezhaFan/translate-cli/internal/model"
)

func TransAndPrint(lines []string, llm model.LLM, wg *sync.WaitGroup) chan string {
	wg.Add(1)
	// 提取需要翻译片段
	beforeTrans := parse(lines)

	// 并发去翻译
	done := make(chan string)
	go func() {
		afterTrans, err := llm.Translate(beforeTrans)
		if err != nil {
			afterTrans = err.Error()
		}
		// 阻塞
		done <- afterTrans
		wg.Done()
	}()

	return done
}

// 解析
func parse(lines []string) (beforeTrans string) {
	var maxLen int
	for _, line := range lines {
		if maxLen < len(line) && len(line) < 60 {
			maxLen = len(line)
		}
	}
	var buff bytes.Buffer
	buff.Grow(2048)
	for _, line := range lines {
		buff.WriteString(line)
		if len(line) < 60 && len(line) < maxLen {
			buff.Write(bytes.Repeat([]byte{' '}, maxLen-len(line)))
		}
		buff.WriteByte('\n')
	}
	return buff.String()
}
