package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := 30 * time.Second

	config := NewConfigFromEnv()
	omeda := NewOmedaClient(config.BaseURL, timeout)
	service := NewOmedaService(omeda)

	hero, err := service.GetHero("Aurora")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", hero)
}
