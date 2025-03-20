package model

import "fmt"

var (
	_ Model = (*Ollama)(nil)
)

type Ollama struct {
	Config
}

func (o *Ollama) Translate(content string) (string, error) {
	var res Response1
	err := post(o.Config, content, &res)
	if err != nil {
		return "", err
	}
	if res.Error != "" {
		return "", fmt.Errorf("%s", res.Error)
	}
	return res.Message.Content, err
}
