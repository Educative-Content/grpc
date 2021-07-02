package main

import (
	"fmt"
	"os"

	_ "github.com/mactsouk/grpc-ch12"
	"github.com/mactsouk/protobuf"
	"google.golang.org/grpc"
)

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
		return
	} else {
		port = os.Args[1]
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	client := protobuf.NewMessageService(conn)

}
