package stdin

import (
	"bufio"
	"os"
)

type StdinReader struct {
	Lines chan []byte
	Err   error
}

func Read() *StdinReader {
	sr := &StdinReader{
		Lines: make(chan []byte),
		Err:   nil,
	}

	go func() {
		defer close(sr.Lines)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			sr.Lines <- scanner.Bytes()
		}
		sr.Err = scanner.Err()
	}()

	return sr
}
