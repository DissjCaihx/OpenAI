package body

import (
	"maps"
)

type RequestBody interface {
	ToMap() map[string]interface{}
	SetOtherParam(map[string]interface{})
	ToParam() map[string]interface{}
	SetHeaders(map[string]interface{})
}
type OpanAiParameters interface {
	Deepseek(string, []map[string]interface{}) Completions
}
type CreateCompletions struct {
}

type Completions struct {
	Messages         []map[string]interface{}
	Model            string
	FrequencyPenalty int
	MaxTokens        int
	Stream           bool
	Temperature      float64
	OtherParam       map[string]interface{}
}

func (c *CreateCompletions) Deepseek(module string, message []map[string]interface{}) Completions {
	return Completions{
		Messages:    message,
		Model:       module,
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
}
func (c *Completions) SetOtherParam(var1 map[string]interface{}) {
	c.OtherParam = var1
}
func (c *Completions) ToMap() map[string]interface{} {
	param := map[string]interface{}{
		"messages":          c.Messages,
		"model":             c.Model,
		"frequency_penalty": c.FrequencyPenalty,
		"max_tokens":        c.MaxTokens,
		"stream":            c.Stream,
		"temperature":       c.Temperature,
	}
	maps.Copy(param, c.OtherParam)
	return param
}
func (c *Completions) ToParam() map[string]interface{} {
	param := map[string]interface{}{}
	maps.Copy(param, c.OtherParam)
	return param
}
func (c *Completions) SetHeaders(headers map[string]interface{}) {}
func (c *Completions) IsEmpty() bool {
	if c.Messages == nil || len(c.Messages) == 0 {
		return true
	}
	if c.Model == "" || len(c.Model) == 0 {
		return true
	}
	return false
}

type HeaderBody struct {
	ContentType   string
	Accept        string
	Authorization string
	Headers       map[string]interface{}
}

func (h *HeaderBody) ToMap() map[string]interface{} {
	param := map[string]interface{}{
		"Content-Type": h.ContentType,
		"Accept":       h.Accept,
		//"authorization": "Bearer " + h.Authorization,
	}
	maps.Copy(param, h.Headers)
	return param
}
func (h *HeaderBody) SetHeaders(var1 map[string]interface{}) {
	h.Headers = var1
	//tools.Copy(h.Headers, var1)
}
func (h *HeaderBody) ToParam() map[string]interface{} {
	param := map[string]interface{}{}
	maps.Copy(param, h.Headers)
	return param
}
func (h *HeaderBody) SetOtherParam(var1 map[string]interface{}) {}
func (h *HeaderBody) IsEmpty() bool {
	if h.ContentType == "" && h.Accept == "" && h.Authorization == "" {
		return true
	}
	return false
}
