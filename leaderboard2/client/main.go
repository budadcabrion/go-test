package main

import (
	"context"
	"log"
	"time"
	"os"
	"strconv"
	"io"

	"github.com/budadcabrion/go-practice/service"

	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
	defaultCmd = "time"
)

func main() {
	// connect to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := service.NewServiceClient(conn)

	// minimal arg parsing
	cmd := defaultCmd
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	// send a message to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch cmd {
	case "time":
		r, err := client.Time(ctx, &service.TimeRequest{})
		if err != nil {
			log.Fatalf("could not get time: %v", err)
		}
		t := time.Unix(r.Timestamp, 0)
		log.Printf("Time: %d, %s", r.Timestamp, t.String())

	case "insert":
		if len(os.Args) < 4 {
			log.Fatalf("not enough args for insert, need 2")
		}
		name := os.Args[2]
		thingType := os.Args[3]
		r, err := client.InsertThing(ctx, &service.Thing{Name: name, Type: thingType})
		if err != nil {
			log.Fatalf("could not insert thing: %v", err)
		}
		log.Printf("Id: %d", r.Id)

	case "get":
		if len(os.Args) < 3 {
			log.Fatalf("not enough args for insert, need 2")
		}
		id, _ := strconv.Atoi(os.Args[2])
		r, err := client.GetThing(ctx, &service.ThingId{Id: int64(id)})
		if err != nil {
			log.Fatalf("could not get thing: %v", err)
		}
		log.Printf("Thing: %v", r)

	case "list":
		stream, err := client.ListThings(ctx, &service.ListThingsRequest{})
		if err != nil {
			log.Fatalf("could not list things: %v", err)
		}

		for {
			thing, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("could not recv thing: %v", err)
			}
			log.Printf("Thing: %v", thing)
		}

	default:
		log.Fatalf("invalid command: %v", cmd)
	}
}