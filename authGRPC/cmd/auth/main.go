package main

import (
	"fmt"
	"github.com/NorthDice/AuthGRPC/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
