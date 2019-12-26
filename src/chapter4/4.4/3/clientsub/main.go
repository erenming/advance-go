package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "advance-go/src/chapter4/4.4/3/pubsubservice"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &pb.String{Value: "golang:"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		fmt.Println(reply.GetValue())
	}
}
