package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	routes.SetRoutes(router)

	err := router.Run(":8080")

	if err != nil {
		panic("Boot process of server failed")
	}
}
