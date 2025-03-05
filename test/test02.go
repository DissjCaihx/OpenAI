package main

import (
	"fmt"
	"github.com/DissjCaihx/OpenAI/body"
	"github.com/DissjCaihx/OpenAI/openai"
)

func main() {
	ai := openai.OpenAI()
	ai.SetApiKey("")
	ai.SetBaseUrl("https://api.deepseek.com")
	headBody := body.HeaderBody{Accept: "application/json", ContentType: "application/json"}
	ai.SetHeaders(headBody)
	createCompletions := body.CreateCompletions{}
	deepseek := createCompletions.Deepseek("deepseek-chat", []map[string]interface{}{
		{
			"content": "You are a helpful assistant",
			"role":    "system",
		},
		{
			"content": "Hi",
			"role":    "user",
		},
	})
	m, err := ai.Completions(deepseek)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
