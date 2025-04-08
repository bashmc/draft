package main

import (
	"log"

	"github.com/topkobie/draft/handlers"
)

func main() {

	srv := AppServer{
		handler: &handlers.ApiHandler{},
	}

	log.Print("starting server..")
	log.Fatal(srv.Start())
}
