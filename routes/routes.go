package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/handlers"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/random/add", handlers.HandleRandomAdd)
	router.GET("/random/sub", handlers.HandleRandomSub)
	router.POST("/add", handlers.HandleAdd)
	router.POST("/sub", handlers.HandleSub)
}
