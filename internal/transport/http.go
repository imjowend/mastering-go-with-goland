package transport

import (
	"encoding/json"
	"github.com/imjowend/mastering-go-with-goland/internal/todo"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, response *http.Request) {
		var t TodoItem
		err := json.NewDecoder(response.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		err = todoSvc.Add(t.Item)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /todo", func(writer http.ResponseWriter, response *http.Request) {

		err := json.NewEncoder(writer).Encode(todoSvc.GetAll())
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	})
	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
