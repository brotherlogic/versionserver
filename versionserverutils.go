package main

import (
	"io/ioutil"

	"github.com/golang/protobuf/proto"

	pb "github.com/brotherlogic/versionserver/proto"
)

func (s *Server) loadVersions() error {
	dirs, _ := ioutil.ReadDir(s.dir)
	for _, f := range dirs {
		data, _ := ioutil.ReadFile(s.dir + "/" + f.Name())
		version := &pb.Version{}
		proto.Unmarshal(data, version)
		s.versions = append(s.versions, version)
	}

	return nil
}

func (s *Server) saveVersions() {
	for _, v := range s.versions {
		data, _ := proto.Marshal(v)
		ioutil.WriteFile(s.dir+"/"+v.GetKey(), data, 0700)
	}
}
