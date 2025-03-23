package model

import (
	"errors"
	"os"
	"strings"
)

type LLM interface {
	Translate(content string) (string, error)
}

type Config struct {
	Type  string
	Url   string
	Model string
	Key   string
	Lang  string
}

func GetLLM() (LLM, error) {
	var conf Config
	conf.Lang = os.Getenv("TC_LANG")
	conf.Url = os.Getenv("TC_LLM_URL")
	conf.Model = os.Getenv("TC_LLM_MODEL")
	conf.Key = os.Getenv("TC_LLM_KEY")
	conf.Type = os.Getenv("TC_LLM_TYPE")
	if conf.Lang == "" {
		conf.Lang = "Chinese"
	}
	if conf.Url == "" || conf.Model == "" || (conf.Key == "" && conf.Type != "ollama") {
		return nil, errors.New("please set environment variables and reload it")
	}
	return newLLM(conf), nil
}

func newLLM(conf Config) LLM {
	typ := strings.TrimSpace(strings.ToLower(conf.Type))
	switch typ {
	case "ollama":
		conf.Url = strings.TrimSuffix(conf.Url, "/") + "/api/chat"
		return &Ollama{conf}
	default:
		conf.Url = strings.TrimSuffix(conf.Url, "/") + "/chat/completions"
		return &Common{conf}
	}
}
