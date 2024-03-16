package main

import (
	grpc_basket "basket-api/grpc"
	"basket-api/repositories"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

func main() {

	server := grpc.NewServer()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	grpc_basket.NewBasketService(server, repositories.NewRedisBasketRepository(log.Default(), client))
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Panicf("Failed to listen: %v", err)
	}

	log.Println("gRPC server started, listening on : 50051")
	if err := server.Serve(lis); err != nil {
		log.Panicf("Failed to serve: %v", err)
	}
}
