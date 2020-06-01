package handler

import (
	"net/http"

	db "../db"
)

func SetUpRouting(postgres *db.Postgres) *http.ServeMux {
	todoHandler := &todoHandler{
		postgres: postgres,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.getAllTodo(w, r)
		case http.MethodPost:
			todoHandler.saveTodo(w, r)
		case http.MethodPut:
			todoHandler.updateTodo(w, r)
		case http.MethodDelete:
			todoHandler.deleteTodo(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})

	return mux
}
