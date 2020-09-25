package main

import (
	"context"
	"log"
	"time"
	"os"
	"strconv"
	"io"

	"github.com/budadcabrion/go-test/leaderboard2/service"

	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
	defaultCmd = ""
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
	case "setscore":
		if len(os.Args) < 4 {
			log.Fatalf("not enough args for setscore, need 2")
		}
		name := os.Args[2]
		score, _ := strconv.ParseInt(os.Args[3], 10, 64)
		_, err := client.SetScore(ctx, &service.PlayerScore{Name: name, Score: score})
		if err != nil {
			log.Fatalf("could not set score: %v", err)
		}
		log.Printf("confirmed %v = %d", name, score)

	case "getscores":
		if len(os.Args) < 4 {
			log.Fatalf("not enough args for setscore, need 2")
		}
		start, _ := strconv.ParseInt(os.Args[2], 10, 64)
		count, _ := strconv.ParseInt(os.Args[3], 10, 64)
		stream, err := client.GetScores(ctx, &service.GetScoresRequest{Start: start, Count: count})
		if err != nil {
			log.Fatalf("could not get scores: %v", err)
		}

		for {
			playerScore, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("could not recv playerscore: %v", err)
			}
			log.Printf("PlayerScore: %v", playerScore)
		}

	default:
		log.Fatalf("invalid command: %v", cmd)
	}
}