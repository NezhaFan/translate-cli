package model

import "fmt"

type Common struct {
	Config
}

func (c *Common) Translate(content string) (string, error) {
	var res Response2
	err := post(c.Config, content, &res)
	if err != nil {
		return "", err
	}
	if res.Error.Message != "" {
		return "", fmt.Errorf("%s", res.Error.Message)
	}
	if len(res.Choices) == 0 {
		return "", fmt.Errorf("返回结构异常: %v", res)
	}
	return res.Choices[0].Message.Content, err
}
