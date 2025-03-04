package main

import (
	"OpenAI/body"
	"OpenAI/openai"
	"fmt"
)

func main() {
	ai := openai.OpenAI()
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
		MaxTokens:   1024,
		Temperature: 0.7,
		Stream:      false,
		OtherParam: map[string]interface{}{
			"frequency_penalty": 0,
			"presence_penalty":  0,
			"response_format": map[string]interface{}{
				"type": "text",
			},
			"stop":           "None",
			"stream_options": "None",
			"temperature":    1,
			"top_p":          1,
			"tools":          "None",
			"tool_choice":    "none",
			"logprobs":       "False",
			"top_logprobs":   "None",
		},
	}
	ai.SetHeaders(headBody)
	fmt.Println(completions)
	//m, err := ai.Completions(completions)
	m, err := ai.Models()
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
