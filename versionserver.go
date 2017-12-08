package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/brotherlogic/goserver"
	"google.golang.org/grpc"

	pbg "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/versionserver/proto"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	versions []*pb.Version
	dir      string
}

// Init builds the server
func Init(dir string) *Server {
	s := &Server{GoServer: &goserver.GoServer{}, dir: dir}

	if _, err := os.Stat(s.dir); os.IsNotExist(err) {
		os.Mkdir(s.dir, 0700)
	}

	s.loadVersions()
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterVersionServerServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Mote promotes/demotes this server
func (s *Server) Mote(master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{&pbg.State{Key: "keys", Value: int64(len(s.versions))}}
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	var dir = flag.String("save_dir", "tmp", "Directory in which to save all the files")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init(*dir)
	server.PrepServer()
	server.Register = server

	server.RegisterServer("versionserver", false)
	server.Log("Starting!")
	server.Serve()
}
