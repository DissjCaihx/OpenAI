package openai

import (
	"github.com/DissjCaihx/OpenAI/body"
	"github.com/DissjCaihx/OpenAI/internal/Client"
	"github.com/DissjCaihx/OpenAI/tools"
	"time"
)

//var (
//	client = Client.Client{}
//)

type openAI struct {
	apiKey     string
	baseUrl    string
	header     body.HeaderBody
	timeSecond time.Duration
	client     Client.Client
}

type OpenAIImpl interface {
	SetTimeSecond(time.Duration)
	SetApiKey(string)
	GetApiKey() string
	SetBaseUrl(string)
	GetBaseUrl() string
	SetHeaders(body.HeaderBody)
	Models() (map[string]interface{}, error)
	Balance() (map[string]interface{}, error)
	Completions(body.Completions) (map[string]interface{}, error)
	CompletionsCreate(string, body.Completions) (map[string]interface{}, error)
}

func (o *openAI) GetApiKey() string {
	return o.apiKey
}
func (o *openAI) GetBaseUrl() string {
	return o.baseUrl
}
func (o *openAI) SetTimeSecond(timeSecond time.Duration) {
	o.timeSecond = timeSecond
	o.client.SetTimeSecond(timeSecond)
}
func (o *openAI) SetHeaders(h body.HeaderBody) {
	o.header = h
	o.client.SetHeaders(h)
}
func (o *openAI) SetApiKey(apiKey string) {
	o.apiKey = apiKey
	o.header.Authorization = "Bearer " + apiKey
	o.client.SetApiKey(apiKey)
}
func (o *openAI) SetBaseUrl(baseUrl string) {
	o.baseUrl = baseUrl
	o.client.SetBaseUrl(baseUrl)
}
func (o *openAI) Models() (map[string]interface{}, error) {
	models, err := o.client.Models()
	if err != nil {
		return nil, err
	}
	return models, nil
}
func (o *openAI) Balance() (map[string]interface{}, error) {
	balance, err := o.client.Balance()
	if err != nil {
		return nil, err
	}
	return balance, nil
}
func (o *openAI) Completions(var1 body.Completions) (map[string]interface{}, error) {
	completions, err := o.client.Completions(var1)
	if err != nil {
		return nil, err
	}
	return completions, nil
}
func (o *openAI) CompletionsCreate(route string, var1 body.Completions) (map[string]interface{}, error) {
	create, err := o.client.CompletionsCreate(route, var1)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func OpenAI() OpenAIImpl {
	ai := openAI{}
	ai.client = Client.Client{}
	config := tools.DefaultConfig()
	ai.client.SetConfig(config)
	ai.client.InitHeaders()
	return &ai
}
