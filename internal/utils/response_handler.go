package utils

import (
	"TanAgah/internal/stringResource"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendSuccessResponse(c *gin.Context, res any, err any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		c.Writer.WriteHeader(http.StatusOK)
		c.JSON(http.StatusOK, gin.H{"error": stringResource.GetStrings().UnknownError(c)})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func SendResponseWithCode(c *gin.Context, res any, err any, httpCode int) {
	c.Writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		c.Writer.WriteHeader(httpCode)
		c.JSON(httpCode, gin.H{"error": stringResource.GetStrings().UnknownError(c)})
	} else {
		c.JSON(httpCode, res)
	}
}

func SendError400Response(c *gin.Context, err string) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusBadRequest)
	c.JSON(http.StatusBadRequest, gin.H{"error": err})
}

func SendError404Response(c *gin.Context, err string) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusNotFound)
	c.JSON(http.StatusNotFound, gin.H{"error": err})
}

func SendError401Response(c *gin.Context, err string) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusUnauthorized)
	c.JSON(http.StatusUnauthorized, gin.H{"error": err})
}

func SendDataError403(c *gin.Context, res any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusForbidden)
	c.JSON(http.StatusForbidden, gin.H{"error": res})
}

func SendDataError500(c *gin.Context, res any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusInternalServerError)
	c.JSON(http.StatusInternalServerError, gin.H{"error": res})
}
