package main

import (
	"github.com/imjowend/mastering-go-with-goland/internal/todo"
	"github.com/imjowend/mastering-go-with-goland/internal/transport"
	"log"
)

func main() {

	svc := todo.NewService()
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
