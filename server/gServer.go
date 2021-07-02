package main

import (
	"fmt"
	"net"
	"os"

	api "github.com/mactsouk/protoapi"
	"google.golang.org/grpc"
)

type MessageServer struct {
}

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
		return
	} else {
		port = os.Args[1]
	}

	server := grpc.NewServer()
	var messageServer MessageServer
	api.RegisterRandomServer(server, messageServer)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving requests...")
	server.Serve(listen)
}
