package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chat "github.com/mukhinfa/chat-server/pkg/chat/v1"
)

const (
	grpcPort = 50051
)

type server struct {
	chat.UnimplementedChatServiceServer
}

func (s *server) Create(_ context.Context, req *chat.CreateRequest) (*chat.CreateResponse, error) {
	log.Println(color.RedString("Create chat request"), fmt.Sprintf("%+v", req))
	return &chat.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(_ context.Context, req *chat.DeleteRequest) (*empty.Empty, error) {
	log.Println(color.RedString("Delete chat request"), fmt.Sprintf("%+v", req))
	return &empty.Empty{}, nil
}

func (s *server) SendMessage(_ context.Context, req *chat.SendMessageRequest) (*empty.Empty, error) {
	log.Println(color.RedString("Send message request"), fmt.Sprintf("%+v", req))
	return &empty.Empty{}, nil
}

func main() {
	log.Println(color.GreenString("Starting chat server"))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chat.RegisterChatServiceServer(s, &server{})

	log.Println(color.GreenString("Server is running on %d", grpcPort))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
