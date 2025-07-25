package main

import (
	"cat-facts/internal/service"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	service.InitCalls()

	end := time.Since(start)
	fmt.Println("Duration:", end.Milliseconds())
}
