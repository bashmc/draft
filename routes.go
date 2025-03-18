package main

import "net/http"

func (s *AppServer) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", s.handler.CreateUserHandler)

	return mux
}
