package main

import (
	"gin.exe2/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet)
	r.Run()
}
