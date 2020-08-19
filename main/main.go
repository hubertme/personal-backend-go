package main

import (
	"fmt"
	"wang.hihubert.personal-backend/gin"
)

func main() {
	fmt.Println("Hello, world!")

	err := gin.SetupGinServer()

	if err != nil {
		panic("PANIC!!!! Failed to setup Gin server!!!!!")
	}
}
