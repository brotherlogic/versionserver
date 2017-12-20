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
	if len(os.Args) == 2 {
		answer, err := registry.GetVersion(context.Background(), &pb.GetVersionRequest{Key: os.Args[1]})
		if err != nil {
			log.Fatalf("Error reading version: %v", err)
		}
		log.Printf("Answer = %v", answer)
	} else {
		val, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error parsing number: %v", err)
		}

		answer, err := registry.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: os.Args[1], Value: int64(val)}})
		if err != nil {
			log.Fatalf("Error reading version: %v", err)
		}
		log.Printf("Answer = %v", answer)
	}
}
