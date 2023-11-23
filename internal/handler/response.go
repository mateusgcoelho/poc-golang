package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	logger.Errorf(err.Error())
	ctx.Header("Content-type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":  message,
	})
}

func sendSuccessResponse(ctx *gin.Context, statusCode int, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data": data,
	})
}