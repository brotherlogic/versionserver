package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/versionserver/proto"
)

func TestWeirdLoad(t *testing.T) {
	s := InitTest("weird_test")

	if len(s.versions) != 3 {
		t.Errorf("Failure to load versions")
	}

	key, err := s.GetVersion(context.Background(), &pb.GetVersionRequest{Key: "github.com.brotherlogic.keystore"})
	if err != nil {
		t.Errorf("Error in getting key")
	}

	if key.GetVersion().GetValue() != 417654 {
		t.Errorf("Bad key returned: %v", key)
	}
}
