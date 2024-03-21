package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	desc "github.com/Lokion13/auth-chat-server/pkg/user_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

// CREATE
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// GET
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return &desc.GetResponse{
		User: &desc.User{
			Id:        req.GetId(),
			Name:      "Its works",
			Email:     gofakeit.Email(),
			Role:      0,
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

// UPDATE
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// DELETE
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf(color.GreenString("server listening at %v", listen.Addr()))

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
