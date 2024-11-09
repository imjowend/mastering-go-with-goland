package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

func main() {

	var todos = make([]string, 0)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, response *http.Request) {
		var todo TodoItem
		err := json.NewDecoder(response.Body).Decode(&todo)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		todos = append(todos, todo.Item)
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /todo", func(writer http.ResponseWriter, response *http.Request) {

		err := json.NewEncoder(writer).Encode(todos)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
