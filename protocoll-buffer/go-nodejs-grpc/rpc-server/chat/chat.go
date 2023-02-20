package chat

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Response, error) {
	log.Println("rpc-server: called SayHello")
	return &Response{Body: fmt.Sprintf("New Message: %s", message.Body)}, nil
}

func (s *Server) GetDetails(ctx context.Context, message *Details) (*Response, error) {
	log.Println("rpc-server: called GetDetails")
	return &Response{Body: fmt.Sprintf("Your name is %s and you're %d years old", message.Name, message.Age)}, nil
}
