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
	completions := body.Completions{
		Messages: []map[string]interface{}{
			{
				"content": "You are a helpful assistant",
				"role":    "system",
			},
			{
				"content": "Hi",
				"role":    "user",
			},
		},
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
	ai.SetHeaders(headBody)
	m, err := ai.Completions(completions)
	//m, err := ai.Models()
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
