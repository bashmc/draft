package main

import (

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	// users
	mux.Post("/users", s.handler.CreateUser)
	mux.Post("/users/verify", s.handler.VerifyUser)
	mux.Get("/users/{id}", s.handler.GetUser)
	mux.Delete("/users/{id}", s.handler.DeleteUser)

	return mux
}
