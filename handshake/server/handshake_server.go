package main

import (
	"context"
	"fmt"
	pb "golang-grpc/handshake"
	"golang-grpc/handshake/model"
	"log"
	"net"

	"github.com/jinzhu/configor"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHandshakeServer
}

func (s *server) Handshake(_ context.Context, request *pb.HandshakeRequest) (*pb.HandshakeReply, error) {
	message := ""

	if request.GetHandshakeStatus() == pb.HandshakeStatus_HANDSHAKE_TYPE_REQUESTED {
		request.HandshakeStatus = pb.HandshakeStatus_HANDSHAKE_TYPE_ACCEPTED
	}

	if request.GetHandshakeStatus() == pb.HandshakeStatus_HANDSHAKE_TYPE_ACCEPTED {
		message = fmt.Sprintf("%s Handshake is Accepted, %v", request.GetUsername(), request.GetHandshakeAt())
	} else {
		message = fmt.Sprintf("%s Handshake is Rejected, %v", request.GetUsername(), request.GetHandshakeAt())
	}
	log.Print(message)

	return &pb.HandshakeReply{Message: message}, nil
}

func main() {
	cfg := model.Config{}
	err := configor.Load(&cfg, "/Users/lukmanhafidz/Documents/Belajar/Golang/Golang-gRPC/config.yml")
	if err != nil {
		log.Fatalf("error when try get config.yml: %s", err.Error())
	}

	// host := fmt.Sprintf(":%d", cfg.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Printf("error when listen to host: %v", err)
		return
	}

	grpcNew := grpc.NewServer()
	pb.RegisterHandshakeServer(grpcNew, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcNew.Serve(lis); err != nil {
		log.Printf("error when serve: %v", err)
		return
	}
}
