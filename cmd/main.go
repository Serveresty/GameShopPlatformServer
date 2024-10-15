package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
