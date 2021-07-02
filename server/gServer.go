package main

import (
	"fmt"
	"os"
)

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
		return
	} else {
		port = os.Args[1]
	}

}
