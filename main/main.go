package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("start!")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Send your name as /:name")
	})

	router.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		f, err := os.Open("secret/prikey.txt")
		var prikey string
		if err == nil {
			b, _ := ioutil.ReadAll(f)
			prikey = string(b)
		} else {
			prikey = ""
		}
		defer f.Close()

		ctx.HTML(http.StatusOK, "index.html", gin.H{"name": name, "prikey": prikey})
	})

	router.POST("/", func(ctx *gin.Context) {
		var reqdata struct {
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}
		var resdata struct {
			Hello string `json:"hello"`
		}
		ctx.BindJSON(&reqdata)
		resdata.Hello = reqdata.Firstname + " " + reqdata.Lastname
		ctx.JSON(http.StatusOK, resdata)
	})

	router.Run()
}
