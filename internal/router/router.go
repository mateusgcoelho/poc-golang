package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/poc-golang/internal/handler"
)

func Initialize() {
	handler.InitializeHandlers()
	
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		initializeUserRoutes(v1)
	} 

	router.Run(":8080")
}