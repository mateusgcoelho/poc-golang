package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/poc-golang/internal/handler"
)

func initializeUserRoutes(r *gin.RouterGroup) {
	router := r.Group("/users")
	{
		router.GET("/", handler.GetUsers)
		router.POST("/", handler.CreateUser)
	}
}