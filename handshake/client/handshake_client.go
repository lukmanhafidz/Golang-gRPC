package main

import (
	"context"
	"fmt"
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
	err := configor.Load(&cfg, "/Users/lukmanhafidz/Documents/Belajar/Golang/Golang-gRPC/config.yml") //load config.yml
	if err != nil {
		log.Fatalf("error when try get config.yml: %s", err.Error())
	}

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials())) //init new grpc client
	defer conn.Close()

	newClient := pb.NewHandshakeClient(conn) //init handshake client

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() //for resource effeciency

	reply, err := newClient.Handshake(ctx, &pb.HandshakeRequest{ //dial handshake server from this client
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
