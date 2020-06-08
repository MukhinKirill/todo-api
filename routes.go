package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpRouting(postgres *Postgres) *mux.Router {
	todoHandler := &todoHandler{
		postgres: postgres,
	}

	router := mux.NewRouter()
	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.getAllTodo(w, r)
		case http.MethodPost:
			todoHandler.saveTodo(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})
	router.HandleFunc("/todo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		switch r.Method {
		case http.MethodGet:
			todoHandler.getTodo(w, r, id)
		case http.MethodPut:
			todoHandler.updateTodo(w, r, id)
		case http.MethodDelete:
			todoHandler.deleteTodo(w, r, id)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})
	http.Handle("/", router)
	return router
}
