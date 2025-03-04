package main

import (
	"OpenAI/body"
	"OpenAI/openai"
	"fmt"
)

func main() {
	ai := openai.OpenAI()
	ai.SetApiKey("sk-49e2ba1915a44ec0aee09c027fa06e0d")
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

	ai.SetHeaders(headBody)
	fmt.Println(deepseek)
	fmt.Println(completions)
	m, err := ai.Completions(deepseek)
	//m, err := ai.Models()
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
