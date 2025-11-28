package main

import (
	"fmt"
	"go-posts/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
