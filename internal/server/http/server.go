package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server interface {
	Serve() error
}

type server struct {
	httpAddr string
	handler  Handler
}

func NewServer(h Handler, httpAddr string) Server {
	return &server{httpAddr: httpAddr, handler: h}
}

func (s *server) Serve() error {

	r := mux.NewRouter()

	r.HandleFunc("/flatten", s.handler.Flatten).Methods(http.MethodPost)
	r.HandleFunc("/history", s.handler.History).Methods(http.MethodGet)

	http.Handle("/", r)
	log.Print(s.httpAddr)

	return http.ListenAndServe(s.httpAddr, nil)
}
