package model

import (
	"strings"
)

type Model interface {
	Translate(content string) (string, error)
}

type Config struct {
	Url   string
	Model string
	Key   string
	Lang  string
}

func New(typ string, conf Config) Model {
	typ = strings.TrimSpace(strings.ToLower(typ))
	switch typ {
	case "ollama":
		conf.Url = strings.TrimSuffix(conf.Url, "/") + "/api/chat"
		return &Ollama{conf}
	default:
		conf.Url = strings.TrimSuffix(conf.Url, "/") + "/chat/completions"
		return &Common{conf}
	}
}
