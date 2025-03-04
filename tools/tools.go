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

// Config 结构体
type Config struct {
	Path              string
	FileName          string
	propertyMap       map[string]interface{}
	propertySourceMap map[string]interface{}
}
type (
	// handlerYamlProperty 接口
	handlerYamlProperty interface {
		get(key string) interface{}
		set(key string, value interface{})
		SetDefault(*Config)
		GetDefault() *Config
		loadProperties() *Config
	}
)

// 默认配置
func DefaultConfig() *Config {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Path:     dir,
		FileName: "openai.yml",
	}
}

// 实现 handlerYamlProperty 接口
// Get 方法
func (c *Config) Get(key string) interface{} {
	return c.propertyMap[key]
}

// Set 方法
func (c *Config) Set(key string, value interface{}) {
	c.propertyMap[key] = value
}

// getDefault 方法
func (c *Config) getDefault() *Config {
	_, err := os.Stat(c.Path + c.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			//config := DefaultConfig()
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

// LoadProperties 方法
func (c *Config) LoadProperties() *Config {
	defaultConfig := c.getDefault()
	if defaultConfig == nil {
		return nil
	}
	log.Printf("load file path:" + defaultConfig.FileName)
	// 加载 YAML 文件
	yamlPath := filepath.Join(defaultConfig.Path, defaultConfig.FileName)
	if err := c.LoadFromYAML(yamlPath); err != nil {
		log.Printf("Failed to load YAML file: %v\n", err)
		return defaultConfig
	}

	return c
}

// LoadFromYAML 从 YAML 文件加载配置
func (c *Config) LoadFromYAML(filePath string) error {
	// 读取文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			//log.Printf("YAML file not found, using default config: %s\n", filePath)
			return nil
		}
		return fmt.Errorf("failed to read YAML file: %w", err)
	}

	// 解析 YAML 到 propertyMap
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
func CopyNotNullValueOf[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		if !reflect.ValueOf(v).IsZero() {
			dst[k] = v
		}
	}

}
