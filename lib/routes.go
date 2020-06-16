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
	router.HandleFunc("/api/todos/", todoHandler.Get).Methods("GET")
	router.HandleFunc("/api/todos/", todoHandler.Create).Methods("POST")
	router.HandleFunc("/api/todos/{id:[0-9]+}", todoHandler.GetById).Methods("GET")
	router.HandleFunc("/api/todos/{id:[0-9]+}", todoHandler.Put).Methods("PUT")
	router.HandleFunc("/api/todos/{id:[0-9]+}", todoHandler.Delete).Methods("DELETE")
	http.Handle("/", router)
	return router
}
