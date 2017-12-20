package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"

	pb "github.com/brotherlogic/versionserver/proto"
)

func (s *Server) loadVersions() error {
	v, err := os.Getwd()
	s.Log(fmt.Sprintf("Loading %v from %v but %v", s.dir, v, err))
	dirs, _ := ioutil.ReadDir(s.dir)
	for _, f := range dirs {
		data, _ := ioutil.ReadFile(s.dir + "/" + f.Name())
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
			return err
		}
	}

	return nil
}
