package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

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

func AskPass(ctx context.Context, m protoapi.RandomClient, seed int64, length int64) (*protoapi.RandomPass, error) {
	request := &protoapi.RequestPass{
		Seed:   seed,
		Length: length,
	}

	r, err := m.GetRandomPass(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func AskRandom(ctx context.Context, m protoapi.RandomClient, seed int64, place int64) (*protoapi.RandomInt, error) {
	request := &protoapi.RandomParams{
		Seed:  seed,
		Place: place,
	}

	r, err := m.GetRandom(ctx, request)
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

	rand.Seed(time.Now().Unix())
	seed := int64(rand.Intn(100))

	client := protoapi.NewRandomClient(conn)
	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Date and Time:", r.Value)

	length := int64(rand.Intn(20))
	p, err := AskPass(context.Background(), client, 100, length+1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Password:", p.Password)

	place := int64(rand.Intn(100))
	i, err := AskRandom(context.Background(), client, seed, place)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 1:", i.Value)

	k, err := AskRandom(context.Background(), client, seed, place-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 2:", k.Value)
}
