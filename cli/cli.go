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
	s, err := registry.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: os.Args[1], Value: 1234}})
	if err != nil {
		log.Fatalf("Error setting version: %v", err)
	}
	log.Printf("Set = %v", s)
	answer, err := registry.GetVersion(context.Background(), &pb.GetVersionRequest{Key: os.Args[1]})
	if err != nil {
		log.Fatalf("Error reading version: %v", err)
	}
	log.Printf("Answer = %v", answer)
}
