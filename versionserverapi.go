package main

import (
	"fmt"

	pb "github.com/brotherlogic/versionserver/proto"
	"golang.org/x/net/context"
)

//GetVersion gets the version for the given key
func (s *Server) GetVersion(ctx context.Context, in *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	for _, v := range s.versions {
		if v.GetKey() == in.GetKey() {
			return &pb.GetVersionResponse{Version: v}, nil
		}
	}

	return nil, fmt.Errorf("Unable to locate key %v", in.GetKey())
}
