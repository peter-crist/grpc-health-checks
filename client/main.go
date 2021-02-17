package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"

	"github.com/peter-crist/grpc-health-checks/protos"
	"google.golang.org/grpc"
)

func main() {
	socketPtr := flag.String("socket", ":5000", "grpc socket")
	flag.Parse()

	conn, err := grpc.Dial(*socketPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ failed to connect: %v", err)
	}
	defer conn.Close()

	log.Println("Launching client to the 🌙 ...")

	client := protos.NewAngryServiceClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		request := &protos.Message{
			Body: message,
		}

		response, err := client.Greeting(context.Background(), request)
		if err != nil {
			log.Printf("🤮 the server doesn't want to talk to you right now: %v", err)
		}

		log.Println(response)
	}
}
