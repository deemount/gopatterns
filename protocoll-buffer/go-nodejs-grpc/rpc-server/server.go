package main

import (
	"fmt"
	"log"
	"net"
	"rpc-server/chat/rpc-server/chat"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("rpc-server: listen on port 9000")

	chatServer := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
	i := 101
	fmt.Println(i)
}
