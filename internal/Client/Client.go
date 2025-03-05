package Client

import (
	"fmt"
	"github.com/DissjCaihx/OpenAI/body"
	"github.com/DissjCaihx/OpenAI/internal/httpClient"
	"github.com/DissjCaihx/OpenAI/tools"
	"time"
)

//var (
//	config = tools.DefaultConfig()
//)

type Client struct {
	ApiKey     string
	BaseUrl    string
	header     map[string]interface{}
	TimeSecond time.Duration
	config     *tools.Config
}

type ClientImpl interface {
	SetTimeSecond(time.Duration)
	SetApiKey(string)
	SetConfig(*tools.Config)
	SetBaseUrl(string)
	InitHeaders()
	GetHeaders() map[string]interface{}
	SetHeaders(body.HeaderBody)
	Models() (map[string]interface{}, error)
	Balance() (map[string]interface{}, error)
	Completions(body.Completions) (map[string]interface{}, error)
	CompletionsCreate(string, body.Completions) (map[string]interface{}, error)
}

func (c *Client) SetConfig(var1 *tools.Config) {
	c.config = var1
}
func (c *Client) SetTimeSecond(timeSecond time.Duration) {
	c.TimeSecond = timeSecond
}
func (c *Client) SetHeaders(h body.HeaderBody) {
	//c.header = h.ToMap()
	//c.header = map[string]interface{}{}
	tools.CopyIsNotNull(c.header, h.ToMap())
}
func (c *Client) GetHeaders() map[string]interface{} {
	return c.header
}
func (c *Client) InitHeaders() {
	c.config.LoadProperties()
	var openai = c.config.Get("openai").(map[string]interface{})
	if c.ApiKey == "" {
		c.ApiKey = fmt.Sprintf("%s", openai["API_KEY"])
	}
	if c.BaseUrl == "" {
		c.BaseUrl = fmt.Sprintf("%s", openai["BASE_URL"])
	}
	if c.TimeSecond == 0 {
		c.TimeSecond = 30 * time.Second
	}
	c.header = make(map[string]interface{})
	c.header["Accept"] = "application/json"
	c.header["Authorization"] = "Bearer " + c.ApiKey
}
func (c *Client) SetApiKey(apiKey string) {
	c.ApiKey = apiKey
	c.header["Authorization"] = "Bearer " + apiKey
}
func (c *Client) SetBaseUrl(baseUrl string) {
	c.BaseUrl = baseUrl
}
func (c *Client) Models() (map[string]interface{}, error) {
	if c.header == nil {
		c.InitHeaders()
	}
	var openai = c.config.Get("openai").(map[string]interface{})
	router := openai["router"].(map[string]interface{})
	httpClient := httpClient.NewHttpClient(c.TimeSecond)
	get, _ := httpClient.GetJSON(c.BaseUrl+fmt.Sprintf("%s", router["models"]), c.header)
	return get, nil
}
func (c *Client) Balance() (map[string]interface{}, error) {
	if c.header == nil {
		c.InitHeaders()
	}
	var openai = c.config.Get("openai").(map[string]interface{})
	router := openai["router"].(map[string]interface{})
	httpClient := httpClient.NewHttpClient(c.TimeSecond)
	get, _ := httpClient.GetJSON(c.BaseUrl+fmt.Sprintf("%s", router["balance"]), c.header)
	return get, nil
}
func (c *Client) Completions(var1 body.Completions) (map[string]interface{}, error) {
	if c.header == nil {
		c.InitHeaders()
	}
	toMap := var1.ToMap()
	var openai = c.config.Get("openai").(map[string]interface{})
	router := openai["router"].(map[string]interface{})
	httpClient := httpClient.NewHttpClient(c.TimeSecond)
	json, _ := httpClient.PostJSON(c.BaseUrl+fmt.Sprintf("%s", router["completions"]), c.header, toMap)
	return json, nil
}
func (c *Client) CompletionsCreate(route string, var1 body.Completions) (map[string]interface{}, error) {
	if c.header == nil {
		c.InitHeaders()
	}
	toMap := var1.ToMap()
	httpClient := httpClient.NewHttpClient(c.TimeSecond)
	json, _ := httpClient.PostJSON(c.BaseUrl+route, c.header, toMap)
	return json, nil
}
