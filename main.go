package main

import (
	"github.com/gin-gonic/gin"
)

var method = "GET"
var path = "/ping"
var port = "5000"

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Handle(method, path, handler)
}

func main() {
	router.Run(":" + port)
}
