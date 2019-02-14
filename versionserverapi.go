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
			v.Setter = in.GetSet().GetSetter()
			found = true
		}
	}

	if !found {
		s.versions = append(s.versions, in.GetSet())
	}

	err := s.saveVersions()
	if err != nil {
		s.Log(fmt.Sprintf("Error writing: %v -> %v", in, err))
		return nil, err
	}
	return &pb.SetVersionResponse{Response: in.GetSet()}, nil
}

// SetIfLessThan sets the version if the value is less than the given
func (s *Server) SetIfLessThan(ctx context.Context, in *pb.SetIfLessThanRequest) (*pb.SetIfLessThanResponse, error) {
	return nil, fmt.Errorf("Not implemented yet")
}
