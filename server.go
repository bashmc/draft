package main

import (
	"net/http"

	"github.com/g88x/draft/handlers"
)

type AppServer struct {
	handler *handlers.ApiHandler
}

func (s *AppServer) Start() error {

	server := http.Server{
		Addr:    ":8080",
		Handler: s.routes(),
	}

	return server.ListenAndServe()
}
