package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/versionserver/proto"
)

func TestFailCWD(t *testing.T) {
	s := InitTest(".testfailcwd")
	s.db = testDiskBridge{failCwd: true}

	err := s.loadVersions()
	if err == nil {
		t.Errorf("Bad cwd has not failed load versions")
	}
}

func TestFailReadDir(t *testing.T) {
	s := InitTest(".testfailreaddir")
	s.db = testDiskBridge{failReadDir: true}

	err := s.loadVersions()
	if err == nil {
		t.Errorf("Bad readdir has not failed load versions")
	}
}

func TestFailReadFile(t *testing.T) {
	s := InitTest(".testfailreadfile")
	s.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: "blah", Value: 1234}})
	s.db = testDiskBridge{failRead: true}

	err := s.loadVersions()
	if err == nil {
		t.Errorf("Bad read file has not failed load versions")
	}
}
