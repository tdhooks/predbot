package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := 30 * time.Second

	config := NewConfigFromEnv()
	omeda := NewOmedaClient(config.BaseURL, timeout)

	items, err := omeda.Get("/items.json")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(items))
}
