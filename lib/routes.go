package todos

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouting(postgres *Postgres) *mux.Router {
	todoHandler := &todoHandler{
		postgres: postgres,
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/todos/", todoHandler.getAllTodo).Methods("GET")
	router.HandleFunc("/api/todos/", todoHandler.saveTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id:[0-9]+}", todoHandler.getTodo).Methods("GET")
	router.HandleFunc("/api/todos/{id:[0-9]+}", todoHandler.updateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id:[0-9]+}", todoHandler.deleteTodo).Methods("DELETE")
	http.Handle("/", router)
	return router
}
