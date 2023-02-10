package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LogErrorWithAbort : Handle error with abort ,
// Parameters: context *gin.Context | err error | status_code int
func LogErrorWithAbort(context *gin.Context, err error, statusCode int) {
	context.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error(), "package": "controllers"})
}

// SendResponse : Send a response with Status OK (200) ,
// Parameters: context *gin.Context | message any
func SendResponse(context *gin.Context, message any) {
	context.JSON(http.StatusOK, gin.H{"message": message})
}
