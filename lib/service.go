package todos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type todoHandler struct {
	postgres *Postgres
}

func (handler *todoHandler) Create(w http.ResponseWriter, r *http.Request) {

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
	todo.NoteDate = time.Now()
	isValid, err := govalidator.ValidateStruct(todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !isValid {
		responseError(w, http.StatusBadRequest, "data is invalid")
		return
	}

	id, err := handler.postgres.Insert(&todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}
func (handler *todoHandler) Put(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
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
	todo.NoteDate = time.Now()
	isValid, err := govalidator.ValidateStruct(todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !isValid {
		responseError(w, http.StatusBadRequest, "data is invalid")
		return
	}

	id, err := handler.postgres.Update(&todo, idStr)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if id == 0 {
		responseError(w, http.StatusNotFound, fmt.Sprintf("todo %s not exist", idStr))
		return
	}
	responseOk(w, nil)
}
func (handler *todoHandler) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	deletedRowsCount, err := handler.postgres.Delete(id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if deletedRowsCount == 0 {
		responseError(w, http.StatusNoContent, fmt.Sprintf("todo %s not exist", id))
		return
	}
	responseOk(w, nil)
}
func (handler *todoHandler) GetById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	todo, err := handler.postgres.GetById(id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if todo == nil {
		responseError(w, http.StatusNoContent, fmt.Sprintf("not fount todo with id:%s", id))
		return
	}
	responseOk(w, *todo)
}
func (handler *todoHandler) Get(w http.ResponseWriter, r *http.Request) {

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
