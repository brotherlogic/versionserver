package main

import (
	"log"
	"os"
	"strconv"

	"github.com/brotherlogic/goserver/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/versionserver/proto"
)

func main() {
	ip, port, err := utils.Resolve("versionserver")
	if err != nil {
		log.Fatalf("Error resolving server: %v", err)
	}
	conn, _ := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	defer conn.Close()

	registry := pb.NewVersionServerClient(conn)
	answer, err := registry.GetVersion(context.Background(), &pb.GetVersionRequest{Key: os.Args[1]})
	if err != nil {
		log.Fatalf("Error reading version: %v", err)
	}
	log.Printf("Answer = %v", answer)
}
