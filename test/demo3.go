package main

import (
	"fmt"
	"github.com/DissjCaihx/OpenAI/body"
	"github.com/DissjCaihx/OpenAI/internal/Client"
	"time"
)

//import (
//	"OpenAI/body"
//	"OpenAI/internal/Client"
//	"fmt"
//	"time"
//)

func main() {
	client := Client.Client{}
	client.InitHeaders()
	client.SetTimeSecond(30 * time.Second)
	client.SetApiKey("apikey")
	client.SetBaseUrl("http://127.0.0.1:8080")
	bodyHead := body.HeaderBody{Accept: "application/json", ContentType: "application/json"}
	client.SetHeaders(bodyHead)
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
	m, err := client.Completions(completions)
	if err != nil {
	}
	fmt.Println(m)
}
