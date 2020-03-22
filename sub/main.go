package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("start sub")

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		log.Println("sub: get!")
		ctx.String(http.StatusOK, "From sub module")
	})

	router.Run()
}
