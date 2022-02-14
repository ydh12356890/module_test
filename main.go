package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UnixNano()
	fmt.Println(now)
	fmt.Println("hello")
}