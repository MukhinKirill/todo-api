package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type todoHandler struct {
	postgres *Postgres
}

func (handler *todoHandler) saveTodo(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todo Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.postgres.Insert(&todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}
func (handler *todoHandler) updateTodo(w http.ResponseWriter, r *http.Request, idStr string) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todo Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := handler.postgres.Update(&todo, idStr)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}
func (handler *todoHandler) deleteTodo(w http.ResponseWriter, r *http.Request, id string) {

	if err := handler.postgres.Delete(id); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
func (handler *todoHandler) getTodo(w http.ResponseWriter, r *http.Request, id string) {

	todoList, err := handler.postgres.GetOne(id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todoList)
}
func (handler *todoHandler) getAllTodo(w http.ResponseWriter, r *http.Request) {

	todoList, err := handler.postgres.GetAll()
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todoList)
}

func responseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
