package main

import (
	"context"
	hs "advance-go/src/chapter4/4.4/1/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)


type HelloServiceImpl struct {
	
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hs.String) (*hs.String, error) {
	reply := &hs.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	hs.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer.Serve(lis)
}