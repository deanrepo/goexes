package handler

import "github.com/gin-gonic/gin"

import "net/http"

func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Found me",
	})
}
