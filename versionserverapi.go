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

// SetVersion sets a given version number
func (s *Server) SetVersion(ctx context.Context, in *pb.SetVersionRequest) (*pb.SetVersionResponse, error) {

	found := false
	for _, v := range s.versions {
		if v.GetKey() == in.GetSet().GetKey() {
			v.Value = in.GetSet().GetValue()
			found = true
		}
	}

	if !found {
	        s.Log(fmt.Sprintf("Writing log: %v -> %v", s.versions, in.GetSet()))
		s.versions = append(s.versions, in.GetSet())
	}

	err := s.saveVersions()
	if err != nil {
		s.Log(fmt.Sprintf("Error writing: %v -> %v", in, err))
		return nil, err
	}
	return &pb.SetVersionResponse{Response: in.GetSet()}, nil
}
