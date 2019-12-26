package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	hs "advance-go/src/chapter4/4.4/1/hello"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := hs.NewHelloServiceClient(conn)
	reply, err:= client.Hello(context.Background(), &hs.String{Value: "hello"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply.GetValue())
}
