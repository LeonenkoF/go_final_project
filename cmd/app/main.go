package main

import (
	"fmt"
	"main/config"
)

func main() {
	cfg := config.New()

	fmt.Println(cfg)
}
