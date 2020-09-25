package main

import (
	"context"
	"log"
	"net"
	"time"

	
	db "github.com/budadcabrion/go-practice/db"
	"github.com/budadcabrion/go-practice/service"

	"google.golang.org/grpc"
)

const (
	port = ":12345"
)

type serviceServer struct {
	service.UnimplementedServiceServer
}

func (s *serviceServer) Time(ctx context.Context, in *service.TimeRequest) (*service.TimeReply, error) {
	log.Printf("Time")
	now := time.Now().Unix()
	return &service.TimeReply{Timestamp: now}, nil
}

func (s *serviceServer) InsertThing(ctx context.Context, thing *service.Thing) (*service.ThingId, error) {
	log.Printf("InsertThing")
	id := db.InsertThing(db.Thing{thing.Id, thing.Name, thing.Type})
	return &service.ThingId{Id: id}, nil
}

func (s *serviceServer) GetThing(ctx context.Context, thingId *service.ThingId) (*service.Thing, error) {
	log.Printf("GetThing")
	thing, err := db.GetThing(thingId.Id)
	return &service.Thing{Id: thing.Id, Name: thing.Name, Type: thing.Type}, err
}

func (s *serviceServer) ListThings(req *service.ListThingsRequest, stream service.Service_ListThingsServer) error {
	log.Printf("ListThings")
	things := db.ListThings()

	for _, thing := range things {
		stream.Send(&service.Thing{Id: thing.Id, Name: thing.Name, Type: thing.Type})
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
