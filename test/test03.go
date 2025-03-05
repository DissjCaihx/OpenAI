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
	completions := body.Completions{
		Messages:    messageBody.ForMessage(),
		Model:       "deepseek-chat",
		MaxTokens:   2048,
		Temperature: 0.7,
		Stream:      false,
		OtherParam: map[string]interface{}{
			"frequency_penalty": 0,
			"presence_penalty":  0,
			"response_format": map[string]interface{}{
				"type": "text",
			},
			"stop":           nil,
			"stream_options": nil,
			"temperature":    1,
			"top_p":          1,
			"tools":          nil,
			"tool_choice":    "none",
			"logprobs":       false,
			"top_logprobs":   nil,
		},
	}
	m, err := ai.Completions(completions)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
