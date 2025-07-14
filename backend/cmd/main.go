package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "url-analyzer-sykell/routes"

)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}
