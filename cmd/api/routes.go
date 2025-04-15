package main

import "net/http"

func (s *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// users
	mux.HandleFunc("POST /users", s.handler.CreateUser)
	mux.HandleFunc("GET /users/{id}", s.handler.GetUser)
	mux.HandleFunc("DELETE /users/{id}", s.handler.DeleteUser)

	return mux
}
