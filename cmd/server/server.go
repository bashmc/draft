package main

import (
	"net/http"

	"github.com/shcmd/draft/handlers"
)

type application struct {
	handler *handlers.Handler
}

func (app *application) start() error {

	server := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	return server.ListenAndServe()
}
