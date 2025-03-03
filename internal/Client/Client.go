package Client

import (
	"OpenAI/internal/httpClient"
	"OpenAI/tools"
	"fmt"
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
	Completions(map[string]string)
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
	router := config.Get("router").(map[string]interface{})
	httpClient := httpClient.HttpClient{}
	get, _ := httpClient.GetJSON(c.BaseUrl+fmt.Sprintf("%s", router["models"]), c.header)
	return get, nil
}

func (c *Client) Balance(var1 map[string]string) {

}
func (c *Client) Completions(var1 map[string]string) {}
