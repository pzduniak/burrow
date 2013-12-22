package burrow

import (
	"github.com/pzduniak/graval"
)

type Config struct {
	HomePath     string
	Authenticate func(string, string) bool
	Port         int
}

type Server struct {
	ftp *graval.FTPServer
}

func NewServer(cfg Config) *Server {
	server := new(Server)

	factory := &PomFTPFactory{cfg.HomePath, cfg.Authenticate}
	server.ftp = graval.NewFTPServer(&graval.FTPServerOpts{Factory: factory, Port: cfg.Port})

	return server
}

func (s *Server) Listen() error {
	err := s.ftp.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
