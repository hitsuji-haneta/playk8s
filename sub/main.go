package main

import (
	"log"
	"net"

	pb "github.com/hitsuji-haneta/playk8s/sub/protocol"
	"github.com/hitsuji-haneta/playk8s/sub/service"
	"google.golang.org/grpc"
)

func main() {
	log.Println("sub: start")

	listenPort, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	greetService := &service.GreetService{}
	pb.RegisterGreetingServer(server, greetService)
	server.Serve(listenPort)
}
