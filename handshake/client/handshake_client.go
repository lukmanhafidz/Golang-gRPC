package main

import (
	"context"
	pb "golang-grpc/handshake"
	"golang-grpc/handshake/model"
	"log"
	"time"

	"github.com/jinzhu/configor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	cfg := model.Config{}
	err := configor.Load(&cfg, "/Users/lukmanhafidz/Documents/Belajar/Golang/Golang-gRPC/config.yml")
	if err != nil {
		log.Fatalf("error when try get config.yml: %s", err.Error())
	}

	// host := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	newClient := pb.NewHandshakeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := newClient.Handshake(ctx, &pb.HandshakeRequest{
		Username:        "User1",
		HandshakeStatus: pb.HandshakeStatus_HANDSHAKE_TYPE_REQUESTED,
		HandshakeAt:     timestamppb.Now(),
	})

	if err != nil {
		log.Println("Handashake Failed: ", err)
		return
	}

	log.Println(reply.GetMessage())
	return
}
