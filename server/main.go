package main

import (
	"log"
	"net"

	"github.com/codergirlsu/chat-app/server/conversation"
	"google.golang.org/grpc"
)

func main() {
	// create an instance of a gRPC server
	grpcServer := grpc.NewServer()
	// register our server implementation as being able to receive
	// the gRPC requests
	conversation.RegisterChatServer(grpcServer, &conversation.ConversationServer{})

	// listen on port 30080
	lis, err := net.Listen("tcp", "localhost:30080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// connect the gRPC server to the port we are listening on
	grpcServer.Serve(lis)
}
