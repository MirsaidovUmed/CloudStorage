package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Address string
	Port    int
	Mux     *mux.Router
	Logger  *logrus.Logger
}

func (s *Server) Run() {
	http.ListenAndServe(fmt.Sprintf("%v:%v", s.Address, s.Port), s.Mux)
}

func NewServer(
	Address string,
	Port int,
	mux *mux.Router,
	logger *logrus.Logger,
) *Server {
	return &Server{
		Address: Address,
		Port:    Port,
		Mux:     mux,
		Logger:  logger,
	}
}
