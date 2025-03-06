package tools

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type config struct {
	Path              string
	FileName          string
	propertyMap       map[string]interface{}
	propertySourceMap map[string]interface{}
}

type HandlerYamlProperty interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	getDefault() *config
	LoadProperties() *config
}

func (c *config) Get(key string) interface{} {
	return c.propertyMap[key]
}

func (c *config) Set(key string, value interface{}) {
	c.propertyMap[key] = value
}

func (c *config) getDefault() *config {
	_, err := os.Stat(c.Path + c.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			var configMap = make(map[string]interface{})
			var routerMap = make(map[string]interface{})
			c.propertyMap = make(map[string]interface{})
			c.propertySourceMap = make(map[string]interface{})
			configMap["API_KEY"] = ""
			configMap["BASE_URL"] = "http://127.0.0.1/"
			routerMap["models"] = "/models"
			routerMap["balance"] = "/user/balance"
			routerMap["completions"] = "/chat/completions"
			configMap["router"] = routerMap
			c.propertyMap["openai"] = configMap
			CopyIsNotNull(c.propertySourceMap, c.propertyMap)
			return c
		}
	}
	return c
}

func (c *config) LoadProperties() *config {
	defaultConfig := c.getDefault()
	if defaultConfig == nil {
		return nil
	}
	log.Printf("load file path:" + defaultConfig.FileName)
	yamlPath := filepath.Join(defaultConfig.Path, defaultConfig.FileName)
	if err := c.LoadFromYAML(yamlPath); err != nil {
		log.Printf("Failed to load YAML file: %v\n", err)
		return defaultConfig
	}

	return c
}

func (c *config) LoadFromYAML(filePath string) error {
	// 读取文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read YAML file: %w", err)
	}

	if err := yaml.Unmarshal(data, &c.propertyMap); err != nil {
		return fmt.Errorf("failed to parse YAML file: %w", err)
	}
	c.propertySourceMap = make(map[string]interface{})
	c.propertySourceMap = c.propertyMap
	return nil
}

func CopyIsNotNull[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		if !reflect.ValueOf(v).IsZero() {
			dst[k] = v
		}
	}

}

// 默认配置
func DefaultConfig() *config {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return &config{
		Path:     dir,
		FileName: "openai.yml",
		propertyMap: map[string]interface{}{
			"openai": map[string]interface{}{
				"router": map[string]interface{}{
					"models":      "/models",
					"balance":     "/user/balance",
					"completions": "/chat/completions",
				},
			},
		},
		propertySourceMap: map[string]interface{}{
			"openai": map[string]interface{}{
				"router": map[string]interface{}{
					"models":      "/models",
					"balance":     "/user/balance",
					"completions": "/chat/completions",
				},
			},
		},
	}
}
