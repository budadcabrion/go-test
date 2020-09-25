package main

import (
	"context"
	"log"
	"net"

	db "github.com/budadcabrion/go-test/leaderboard2/db"
	"github.com/budadcabrion/go-test/leaderboard2/service"

	"google.golang.org/grpc"
)

const (
	port = ":12345"
)

type serviceServer struct {
	service.UnimplementedServiceServer
}

func (s *serviceServer) SetScore(ctx context.Context, in *service.PlayerScore) (*service.SetScoreReply, error) {
	log.Printf("SetScore %v = %v", in.Name, in.Score)
	db.SetScore(in.Name, in.Score)
	return &service.SetScoreReply{}, nil
}

func (s *serviceServer) GetScores(req *service.GetScoresRequest, stream service.Service_GetScoresServer) error {
	log.Printf("GetScores %v %v", req.Start, req.Count)
	scores := db.GetScores(req.Start, req.Count)

	for _, s := range scores {
		stream.Send(&service.PlayerScore{Name: s.Name, Score: s.Score})
	}
	return nil
}

func main() {
	log.Printf("Starting DB")
	db.InitDB("db", false)
	log.Printf("Starting server")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	service.RegisterServiceServer(grpcServer, &serviceServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to start grpc: %v", err)
	}
}
