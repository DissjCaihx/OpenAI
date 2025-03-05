# OpenAI GoLang

This is a golang third party open source library.Provide a user-friendly API for the operation of large AI models.

### We support
* OpenAi
* Deepseek

More APIs are currently open
## Installation
```
go get github.com/DissjCaihx/OpenAI
```
Currently, OpenAI requires Go version 1.23 or greater.

## Usage
### example 01
* type Client struct{}
```go
import (
"github.com/DissjCaihx/OpenAI/internal/Client"
)

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
```

### example 02
* func OpenAI() OpenAIImpl()
* type OpenAIImpl interface{}
### 02-1
```go
import (
    "fmt"
    "github.com/DissjCaihx/OpenAI/body"
    "github.com/DissjCaihx/OpenAI/openai"
)
func main() {
    ai := openai.OpenAI()
    ai.SetApiKey("apiKey")
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
    if err != nil {
    panic(err)
    }
    fmt.Println(m)
}
```
### 02-2
```go
import (
    "fmt"
    "github.com/DissjCaihx/OpenAI/body"
    "github.com/DissjCaihx/OpenAI/openai"
)
func main() {
    ai := openai.OpenAI()
    ai.SetApiKey("ApiKey")
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
```
### 02-3
```go
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
```
### 02-4
```go
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
    m, err := ai.Completions(deepseek)
    if err != nil {
    panic(err)
    }
    fmt.Println(m)
}
```
## example 03
* Create an open.xml file in the root directory of the project
[![Create an open.xml file]([/Volumes/caihx/projects/OpenAI/test/img.png](https://github.com/DissjCaihx/OpenAI/blob/master/test/img.png))]
```yaml
openai:
  API_KEY: sk-
  BASE_URL: https://api.deepseek.com
  router:
    models: /models
    balance: /user/balance
    completions: /chat/completions
```
```go
func main() {
    ai := openai.OpenAI()
    //ai.SetApiKey("") Automatically read openy.yml ApiKey
    //ai.SetBaseUrl("https://api.deepseek.com") Automatically read openy.yml BaseUrl
    headBody := body.HeaderBody{Accept: "application/json", ContentType: "application/json"}
    ai.SetHeaders(headBody)
    messageBody := body.MessageBody{}
    messageBody.Push(body.Message{Content: "You are a helpful assistant", Role: "system"},
    body.Message{Content: "Hi", Role: "user"})
    createCompletions := body.CreateCompletions{}
    deepseek := createCompletions.Deepseek("deepseek-chat", messageBody.ForMessage())
    m, err := ai.Completions(deepseek)
    if err != nil {
    panic(err)
    }
    fmt.Println(m)
}
```
## example 04
* func (o *openAI) CompletionsCreate(route string, var1 body.Completions) (map[string]interface{}, error) {}
```go
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
    m, err := ai.CompletionsCreate("/chat/completions",deepseek)
    if err != nil {
    panic(err)
    }
    fmt.Println(m)
}
```
## Thank you
