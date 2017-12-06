package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/versionserver/proto"
)

func TestRestart(t *testing.T) {
	s := Init(".testrestart")
	_, err := s.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: "donkey", Value: 1234}})

	s2 := Init(".testrestart")
	val, err := s2.GetVersion(context.Background(), &pb.GetVersionRequest{Key: "donkey"})
	if err != nil {
		t.Fatalf("Error in get version: %v", err)
	}
	if val.GetVersion().GetValue() != 1234 && val.GetVersion().GetKey() != "donkey" {
		t.Errorf("Bad version returned: %v", val)
	}
}

func TestPass(t *testing.T) {
	s := Init(".testpass")
	s.versions = append(s.versions, &pb.Version{Key: "donkey", Value: 1234})
	val, err := s.GetVersion(context.Background(), &pb.GetVersionRequest{Key: "donkey"})
	if err != nil {
		t.Fatalf("Error in get version: %v", err)
	}
	if val.GetVersion().GetValue() != 1234 && val.GetVersion().GetKey() != "donkey" {
		t.Errorf("Bad version returned: %v", val)
	}
}

func TestGetFail(t *testing.T) {
	s := Init(".testfail")
	s.versions = append(s.versions, &pb.Version{Key: "donkey", Value: 1234})
	val, err := s.GetVersion(context.Background(), &pb.GetVersionRequest{Key: "magic"})
	if err == nil {
		t.Fatalf("No error returned?: %v", val)
	}
}

func TestGetWriteFail(t *testing.T) {
	s := Init(".testwritefail/")
	val, err := s.SetVersion(context.Background(), &pb.SetVersionRequest{&pb.Version{Key: "magic/donkey", Value: 1234}})
	if err == nil {
		t.Fatalf("No error returned?: %v", val)
	}
}

func TestSetAndGet(t *testing.T) {
	s := Init(".testsetandget")
	_, err := s.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: "donkey", Value: 1234}})
	if err != nil {
		t.Fatalf("Error in set version: %v", err)
	}
	val, err := s.GetVersion(context.Background(), &pb.GetVersionRequest{Key: "donkey"})
	if err != nil {
		t.Fatalf("Error in get version: %v", err)
	}
	if val.GetVersion().GetValue() != 1234 && val.GetVersion().GetKey() != "donkey" {
		t.Errorf("Bad version returned: %v", val)
	}
}
