package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response1 struct {
	Error   string `json:"error"`
	Message `json:"message"`
}

type Response2 struct {
	Error struct {
		Message string `json:"message"`
	}
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

const (
	promptSystem = "You are a professional programmer and translator, Treat next line as plain text input and translate it into %s,strictly keep \\n symbol and output translation ONLY. If translation is unnecessary (e.g. proper nouns, codes, etc.), return the original text. NO explanations. NO notes."
	// promptAttach = "%s"
)

func post(c Config, content string, to any) error {
	request := &Request{
		Model: c.Model,
		Messages: []Message{
			{Role: "system", Content: fmt.Sprintf(promptSystem, c.Lang)},
			{Role: "user", Content: content},
		},
		Temperature: 0,
	}
	b, _ := json.Marshal(request)
	// fmt.Println(string(b))
	req, err := http.NewRequest(http.MethodPost, c.Url, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	if c.Key != "" {
		req.Header.Add("Authorization", "Bearer "+c.Key)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, _ = io.ReadAll(res.Body)
	// fmt.Printf("返回：：：：%s\n", string(b))
	return json.Unmarshal(b, to)
}
