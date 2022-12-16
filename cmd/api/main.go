package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luanadantaas/roman-numbers-api-go/src/service"
)

func main() {
	r := gin.Default()
	r.POST("/search", service.FindPrimeNum)

	r.Run() // listen and serve on 0.0.0.0:8080

}
