package Client

import (
	"OpenAI/internal/httpClient"
	"OpenAI/openai"
	"OpenAI/tools"
	"fmt"
	"time"
)

var (
	config = tools.DefaultConfig()
)

type Client struct {
	ApiKey  string
	BaseUrl string
	header  map[string]string
}

type ClientImpl interface {
	Client() *Client
	SetApiKey(string)
	SetBaseUrl(string)
	InitHeaders()
	GetHeaders() map[string]string
	Models(map[string]string) (map[string]interface{}, error)
	Balance(map[string]string)
	//Completions(map[string]string)
	Completions(c openai.Completions) (map[string]interface{}, error)
}

func (c *Client) Client() *Client {
	return &Client{}
}
func (c *Client) GetHeaders() map[string]string {
	return c.header
}
func (c *Client) InitHeaders() {
	config.LoadProperties()
	var openai = config.Get("openai").(map[string]interface{})
	if c.ApiKey == "" {
		c.ApiKey = fmt.Sprintf("%s", openai["API_KEY"])
	}
	if c.BaseUrl == "" {
		c.BaseUrl = fmt.Sprintf("%s", openai["BASE_URL"])
	}
	c.header = make(map[string]string)
	c.header["Accept"] = "application/json"
	c.header["Authorization"] = "Bearer " + c.ApiKey
}
func (c *Client) SetApiKey(apiKey string) {
	c.ApiKey = apiKey
}
func (c *Client) SetBaseUrl(baseUrl string) {
	c.BaseUrl = baseUrl
}
func (c *Client) Models(var1 map[string]string) (map[string]interface{}, error) {
	if c.header == nil {
		c.InitHeaders()
	}
	var openai = config.Get("openai").(map[string]interface{})
	router := openai["router"].(map[string]interface{})
	httpClient := httpClient.NewHttpClient(10 * time.Second)
	get, _ := httpClient.GetJSON(c.BaseUrl+fmt.Sprintf("%s", router["models"]), c.header)
	return get, nil
}

func (c *Client) Balance(var1 map[string]string) {

}
func (c *Client) Completions(var1 *openai.Completions) (map[string]interface{}, error) {
	if c.header == nil {
		c.InitHeaders()
	}
	toMap := var1.ToMap()
	var openai = config.Get("openai").(map[string]interface{})
	router := openai["router"].(map[string]interface{})
	httpClient := httpClient.NewHttpClient(10 * time.Second)
	json, _ := httpClient.PostJSON(c.BaseUrl+fmt.Sprintf("%s", router["completions"]), c.header, toMap)
	return json, nil
}
