package main

import (
	"desafio/src/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/search", service.FindPrimeNum)

	r.Run() // listen and serve on 0.0.0.0:8080

}
