package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/brotherlogic/goserver"
	"google.golang.org/protobuf/proto"
	"google.golang.org/grpc"

	pbg "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/versionserver/proto"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	versions   []*pb.Version
	dir        string
	db         diskBridge
	slowDown   bool
	writeMutex *sync.Mutex
}

type prodDiskBridge struct{}

func (p prodDiskBridge) getwd() (string, error) {
	return os.Getwd()
}

func (p prodDiskBridge) readdir(dir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}

func (p prodDiskBridge) read(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

// Init builds the server
func Init(dir string) *Server {
	s := &Server{
		GoServer:   &goserver.GoServer{},
		dir:        dir,
		slowDown:   false,
		writeMutex: &sync.Mutex{},
	}
	s.PrepServer()
	s.Register = s
	s.db = prodDiskBridge{}

	if _, err := os.Stat(s.dir); os.IsNotExist(err) {
		os.Mkdir(s.dir, 0700)
	}

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

// Shutdown shutsdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return s.loadVersions()
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	keys := make([]string, 0)
	size := int64(0)
	for _, k := range s.versions {
		keys = append(keys, k.GetKey())
		size += int64(proto.Size(k))
	}
	return []*pbg.State{
		&pbg.State{Key: "sizington", Value: size},
		&pbg.State{Key: "keys", Text: fmt.Sprintf("%v", keys)},
		&pbg.State{Key: "dir", Text: fmt.Sprintf("%v", s.dir)},
	}
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	var dir = flag.String("save_dir", "/media/scratch/versionserver/", "Directory in which to save all the files")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init(*dir)

	err := server.RegisterServerV2("versionserver", false, false)
	if err != nil {
		return
	}

	err = server.Serve()
	if err != nil {
		fmt.Printf("Error serving: %v\n", err)
	}
}
