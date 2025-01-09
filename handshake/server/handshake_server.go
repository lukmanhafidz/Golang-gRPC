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

type server struct { //UnimplementedHandshakeServer must be embedded for consistency
	pb.UnimplementedHandshakeServer
}

func (s *server) Handshake(_ context.Context, request *pb.HandshakeRequest) (*pb.HandshakeReply, error) { //handshake usecase
	message := ""

	if request.GetHandshakeStatus() == pb.HandshakeStatus_HANDSHAKE_TYPE_REQUESTED {
		request.HandshakeStatus = pb.HandshakeStatus_HANDSHAKE_TYPE_ACCEPTED
	} else {
		request.HandshakeStatus = pb.HandshakeStatus_HANDSHAKE_TYPE_REJECTED
	}

	if request.GetHandshakeStatus() == pb.HandshakeStatus_HANDSHAKE_TYPE_ACCEPTED {
		message = fmt.Sprintf("%s Handshake is Accepted, %v", request.GetUsername(), request.GetHandshakeAt())
	} else {
		message = fmt.Sprintf("%s Handshake is Rejected, %v", request.GetUsername(), request.GetHandshakeAt())
	}
	log.Print(message)

	return &pb.HandshakeReply{Message: message}, nil //reply to client
}

func main() {
	cfg := model.Config{}
	err := configor.Load(&cfg, "/Users/lukmanhafidz/Documents/Belajar/Golang/Golang-gRPC/config.yml") //load config from config.yml
	if err != nil {
		log.Fatalf("error when try get config.yml: %s", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Printf("error when listen to host: %v", err)
		return
	}
	log.Printf("server listening at %v", lis.Addr())

	grpcNew := grpc.NewServer() //initiate new grpc server
	pb.RegisterHandshakeServer(grpcNew, &server{})

	if err := grpcNew.Serve(lis); err != nil { //run new grpc server
		log.Printf("error when serve: %v", err)
		return
	}
}
