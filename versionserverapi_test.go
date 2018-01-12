package main

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	pb "github.com/brotherlogic/versionserver/proto"
)

type testDiskBridge struct {
	failCwd     bool
	failReadDir bool
	failRead    bool
}

func (p testDiskBridge) getwd() (string, error) {
	if p.failCwd {
		return "", errors.New("Designed to fail")
	}
	return os.Getwd()
}

func (p testDiskBridge) readdir(dir string) ([]os.FileInfo, error) {
	if p.failReadDir {
		return make([]os.FileInfo, 0), errors.New("Designed to fail")
	}
	return ioutil.ReadDir(dir)
}

func (p testDiskBridge) read(file string) ([]byte, error) {
	if p.failRead {
		return make([]byte, 0), errors.New("Designed to fail")
	}
	return ioutil.ReadFile(file)
}

func InitTest(dir string) *Server {
	s := Init(dir)
	s.SkipLog = true
	s.db = testDiskBridge{}
	s.loadVersions()
	return s
}

func TestRestart(t *testing.T) {
	s := InitTest(".testrestart")
	s.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: "donkey.magic", Value: 1234}})

	s2 := InitTest(".testrestart")
	val, err := s2.GetVersion(context.Background(), &pb.GetVersionRequest{Key: "donkey.magic"})
	if err != nil {
		t.Fatalf("Error in get version: %v", err)
	}
	if val.GetVersion().GetValue() != 1234 && val.GetVersion().GetKey() != "donkey.magic" {
		t.Errorf("Bad version returned: %v", val)
	}
}

func TestMultiWrite(t *testing.T) {
	s := InitTest(".testmultiwrite")
	s.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: "donkey", Value: 1234}})
	s.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: "donkey", Value: 12345}})

	if len(s.versions) != 1 {
		t.Errorf("Too Many versions: %v", s.versions)
	}
}

func TestPass(t *testing.T) {
	s := InitTest(".testpass")
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
	s := InitTest(".testfail")
	s.versions = append(s.versions, &pb.Version{Key: "donkey", Value: 1234})
	val, err := s.GetVersion(context.Background(), &pb.GetVersionRequest{Key: "magic"})
	if err == nil {
		t.Fatalf("No error returned?: %v", val)
	}
}

func TestGetWriteFail(t *testing.T) {
	s := InitTest(".testwritefail/")
	val, err := s.SetVersion(context.Background(), &pb.SetVersionRequest{Set: &pb.Version{Key: "magic/donkey", Value: 1234}})
	if err == nil {
		t.Fatalf("No error returned?: %v", val)
	}
}

func TestSetAndGet(t *testing.T) {
	s := InitTest(".testsetandget")
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
