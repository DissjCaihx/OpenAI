package main

import (
	"OpenAI/internal/Client"
	"fmt"
)

func main() {
	client := Client.Client{}
	client.SetApiKey("adasjhdajkdnaks")
	headers := client.GetHeaders()
	fmt.Println(headers)
}
