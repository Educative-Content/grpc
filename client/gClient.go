package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mactsouk/protoapi"
	"google.golang.org/grpc"
)

var port = ":8080"

func AskingDateTime(ctx context.Context, m protoapi.RandomClient) (*protoapi.DateTime, error) {
	request := &protoapi.RequestDateTime{
		Value: "Please send me the date and time",
	}

	r, err := m.GetDate(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	client := protoapi.NewRandomClient(conn)
	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Text:", r.Value)
}
