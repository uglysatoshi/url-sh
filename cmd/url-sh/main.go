package main

import (
    "fmt"
    "url-sh/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}