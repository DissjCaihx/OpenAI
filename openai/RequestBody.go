package openai

import "maps"

type Completions struct {
	Messages         []map[string]interface{}
	Model            string
	FrequencyPenalty int
	MaxTokens        int
	Stream           bool
	Temperature      float64
	OtherParam       map[string]interface{}
}
type RequestBody interface {
	ToMap() map[string]interface{}
	SetOtherParam(map[string]interface{})
	ToParam() map[string]interface{}
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
