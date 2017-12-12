package main

import (
	"fmt"
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
	_, err = registry.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: os.Args[1], Value: 1234}})
	if err != nil {
		log.Fatalf("Error setting version: %v", err)
	}
	answer, err := registry.GetVersion(context.Background(), &pb.GetVersionRequest{Key: os.Args[1]})
	if err != nil {
		log.Fatalf("Error reading version: %v", err)
	}
	fmt.Printf("Answer = %v\n", answer)
}
