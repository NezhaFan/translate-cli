//go:generate go build -o tc
package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/NezhaFan/translate-cli/internal/model"
	"github.com/NezhaFan/translate-cli/internal/stdin"
	"github.com/NezhaFan/translate-cli/internal/translator"
)

func main() {

	// for i := 1; i <= 30; i++ {
	// 	fmt.Println("Hello, my number is " + strconv.Itoa(i))
	// 	if i == 4 {
	// 		time.Sleep(time.Second)
	// 	}
	// 	if i == 1 || i == 5 || i == 9 {
	// 		time.Sleep(time.Second * 1)
	// 	}
	// }
	// return

	// 大模型
	llm, err := model.GetLLM()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 读取控制台输出
	sr := stdin.Read()

	// 顺序阻塞结果打印，并且控制并发数
	pipes := make(chan chan string, 5)
	defer close(pipes)
	go func() {
		for ch := range pipes {
			fmt.Print(<-ch)
			close(ch)
		}
	}()

	// 控制程序在打印所有请求完成后才退出
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		// 控制读取结束
		timer := time.NewTicker(time.Millisecond * 300)
		var lines [][]byte

	Loop:
		for {
			select {
			// 逐行读取，合并发送
			case line, ok := <-sr.Lines:
				// 读取完毕
				if !ok {
					if sr.Err != nil {
						fmt.Printf("Error reading from stdin: %v\n", sr.Err)
						os.Exit(1)
					}
					break Loop
				}
				lines = append(lines, line)
			// 超时发送一次
			case <-timer.C:
				if len(lines) > 0 {
					pipes <- translator.TransAndPrint(lines, llm, wg)
					lines = lines[:0]
					timer.Stop()
					timer = time.NewTicker(time.Millisecond * 300) // 重新计时
				}
			}
		}
		if len(lines) > 0 {
			pipes <- translator.TransAndPrint(lines, llm, wg)
		}
	}()

	wg.Wait()
	// 防止wg直接退出导致的最终结果打印不出来问题
	time.Sleep(time.Millisecond * 10)
}
