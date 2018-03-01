package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"

	pb "github.com/brotherlogic/versionserver/proto"
)

type diskBridge interface {
	getwd() (string, error)
	readdir(dir string) ([]os.FileInfo, error)
	read(file string) ([]byte, error)
}

func (s *Server) loadVersions() error {
    	s.versions = make([]*pb.Version, 0)
	_, err := s.db.getwd()
	if err != nil {
		return err
	}
	dirs, err := s.db.readdir(s.dir)
	if err != nil {
		return err
	}
	for _, f := range dirs {
		data, err := s.db.read(s.dir + "/" + f.Name())
		if err != nil {
			return err
		}
		version := &pb.Version{}
		proto.Unmarshal(data, version)
		s.versions = append(s.versions, version)
	}

	return nil
}

func (s *Server) saveVersions() error {
	for _, v := range s.versions {
		data, _ := proto.Marshal(v)
		err := ioutil.WriteFile(s.dir+"/"+v.GetKey(), data, 0700)
		if err != nil {
			return fmt.Errorf("Error writing '%v' - leads to %v", v, err)
		}
	}

	return nil
}
