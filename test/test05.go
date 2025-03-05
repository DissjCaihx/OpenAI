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
	messageBody := body.MessageBody{}
	messageBody.Push(body.Message{Content: "You are a helpful assistant", Role: "system"},
		body.Message{Content: "Hi", Role: "user"})
	createCompletions := body.CreateCompletions{}
	deepseek := createCompletions.Deepseek("deepseek-chat", messageBody.ForMessage())
	m, err := ai.CompletionsCreate("/chat/completions", deepseek)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
