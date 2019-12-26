package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

type MyServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":9000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &MyServer{s}
}

func (s *MyServer) Run() {
	log.Fatal(s.server.ListenAndServe())
}
