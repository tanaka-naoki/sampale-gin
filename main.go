package main

import (
	"github.com/gin-gonic/gin"
	"sample-gin/handler"
	"sample-gin/transecode"
)

func main() {
	transecode := transecode.New()
	r := gin.Default()
	r.GET("/transecode", handler.TransecodeGet(transecode))
	r.POST("/transecode", handler.TransecodePost(transecode))

	r.Run(":8080")
}
