package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/wisaitas/solved-backend-challenge/three/proto"

	"github.com/wisaitas/solved-backend-challenge/one"
	"github.com/wisaitas/solved-backend-challenge/three"
	"github.com/wisaitas/solved-backend-challenge/two"
	"google.golang.org/grpc"
)

func main() {
	// findMaxPath() // can run go test ./...
	// catchMe() // can run go test ./...
	// pireFireDire() // import proto then call http://localhost:50051/beef/summary
}

func findMaxPath() {
	triangle := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	fmt.Println("Solution 1 FindMaxPath :", one.FindMaxPath(triangle))
}

func catchMe() {
	var input string
	fmt.Scan(&input)
	fmt.Println("Solution 2 CatchMe :", two.CatchMe(input))
}

func pireFireDire() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBeefCounterServer(s, &three.BeefServer{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
