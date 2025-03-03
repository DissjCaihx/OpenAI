package main

import (
	"OpenAI/internal/Client"
	"OpenAI/openai"
	"fmt"
)

func main() {
	client := Client.Client{}
	client.InitHeaders()
	models, err := client.Models(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(models)

	completions := openai.Completions{
		Messages: []map[string]interface{}{
			{"role": "user", "content": "Hello, world!"},
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
	m, err := client.Completions(&completions)
	if err != nil {
	}
	fmt.Println(m)
}
