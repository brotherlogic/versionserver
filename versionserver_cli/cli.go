package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/versionserver/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&utils.DiscoveryClientResolverBuilder{})
}

func main() {
	conn, err := grpc.Dial("discovery:///versionserver", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()

	registry := pb.NewVersionServerClient(conn)
	switch os.Args[1] {
	case "get":
		answer, err := registry.GetVersion(context.Background(), &pb.GetVersionRequest{Key: os.Args[2]})
		if err != nil {
			log.Fatalf("Error reading version: %v", err)
		}
		fmt.Printf("Answer = %v [%v]\n", answer, time.Unix(answer.Version.Value, 0))
	case "set":
		val, _ := strconv.Atoi(os.Args[3])
		answer, err := registry.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: os.Args[2], Value: int64(val)}})
		if err != nil {
			log.Fatalf("Error writing version: %v", err)
		}
		fmt.Printf("Answer = %v\n", answer)
	case "guard":
		answer, err := registry.SetIfLessThan(context.Background(),
			&pb.SetIfLessThanRequest{
				Set:          &pb.Version{Key: "guard" + os.Args[2], Value: time.Now().Add(time.Minute * 5).Unix(), Setter: "cli"},
				TriggerValue: time.Now().Unix(),
			})
		if err != nil {
			log.Fatalf("Error reading version: %v", err)
		}
		if !answer.Success {
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %v", os.Args[1])
	}
}
