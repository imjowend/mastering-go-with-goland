package main

import (
	"github.com/imjowend/mastering-go-with-goland/internal/db"
	"github.com/imjowend/mastering-go-with-goland/internal/todo"
	"github.com/imjowend/mastering-go-with-goland/internal/transport"
	"log"
)

// Command + K to see changes to commit

func main() {

	d, err := db.New("postgres", "example", "localhost", "postgres", 5432)
	if err != nil {
		log.Fatal(err)
	}
	svc := todo.NewService(d)
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
