package main

import "fmt"
import "OpenAI/tools"

func main() {
	fmt.Println("Hello World")
	//newConfig := &tools.Config{
	//	Path:     "D:/GoProject/OpenAI/",
	//	FileName: "openai.yml",
	//}

	config := tools.DefaultConfig()
	//newConfig.LoadProperties()
	//fmt.Println(*newConfig)

	config.LoadProperties()
	fmt.Println(*config)
}
