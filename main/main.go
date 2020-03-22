package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	pb "github.com/hitsuji-haneta/playk8s/main/protocol"
	"google.golang.org/grpc"
)

func main() {
	log.Println("main: start!")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Send your name as /:name")
	})

	router.GET("/name/:name", func(ctx *gin.Context) {
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

	router.GET("/push", func(ctx *gin.Context) {
		log.Println("main: push")

		conn, err := grpc.Dial("playk8s-sub-service:8081", grpc.WithInsecure())
		if err != nil {
			log.Fatal("client connection error:", err)
		}
		defer conn.Close()

		client := pb.NewGreetingClient(conn)
		greeter := &pb.Greeter{Name: "tama"}
		res, err := client.Hello(ctx, greeter)
		if err != nil {
			log.Fatal(err)
		}

		ctx.String(http.StatusOK, fmt.Sprintf("%s, %s!", res.GetGreeting(), res.GetName()))
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
