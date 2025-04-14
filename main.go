package main

import (
	"log"

	"github.com/primekobie/draft/handlers"
)

func main() {

	srv := AppServer{
		handler: &handlers.ApiHandler{},
	}

	log.Print("starting server..")
	log.Fatal(srv.Start())
}
