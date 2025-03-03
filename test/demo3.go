package main

import (
	"OpenAI/internal/Client"
	"fmt"
)

func main() {
	client := Client.Client{}
	client.InitHeaders()
	models, err := client.Models(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(models)
}
