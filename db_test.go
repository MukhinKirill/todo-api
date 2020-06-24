package main

import (
	"strconv"
	"testing"
	"time"

	todos "github.com/MukhinKirill/todo-api/lib"
)

func TestDbGetAll(t *testing.T) {
	var postgres *todos.Postgres
	var errDb error
	postgres, errDb = todos.ConnectDb("user=tester password=test host=idm.centos.myidm.local port=5432 dbname=testdb sslmode=disable")
	if errDb != nil {
		panic(errDb)
	} else if postgres == nil {
		panic("postgres is nil")
	}
	todos, err := postgres.GetAll()
	if err != nil {
		t.Errorf("Test DbGetAll() %s", err)
	}
	if len(todos) == 0 {
		t.Errorf("Test DbGetAll() failed, todos is empty")
	}
}

func TestDbInsertUpdateTodoAndGetByIdAndDelete(t *testing.T) {
	var postgres *todos.Postgres
	var errDb error
	postgres, errDb = todos.ConnectDb("user=tester password=test host=idm.centos.myidm.local port=5432 dbname=testdb sslmode=disable")
	if errDb != nil {
		panic(errDb)
	} else if postgres == nil {
		panic("postgres is nil")
	}
	newTodo := todos.Todo{
		Note:     "do test",
		NoteDate: time.Now(),
		Title:    "Test",
	}
	id, err := postgres.Insert(&newTodo)
	if err != nil {
		t.Errorf("Test DbInsert() failed, %s", err)
	}
	todo, err := postgres.GetById(strconv.Itoa(id))
	if err != nil {
		t.Errorf("Test DbGetById() failed, %s", err)
	}
	if todo.ID != id || todo.Note != newTodo.Note || todo.Title != newTodo.Title || todo.NoteDate.Equal(newTodo.NoteDate) {
		t.Errorf("Test DbGetById() failed, gotten note and  created note is not equal")
	}
	var updateTodo todos.Todo
	updateTodo.ID = todo.ID
	updateTodo.Note = "do  update test"
	updateTodo.NoteDate = time.Now()
	updateTodo.Title = "Test update"

	_, err = postgres.Update(&updateTodo, strconv.Itoa(updateTodo.ID))
	if err != nil {
		t.Errorf("Test DbInsert() failed, %s", err)
	}
	todoU, err := postgres.GetById(strconv.Itoa(updateTodo.ID))
	if err != nil {
		t.Errorf("Test DbGetById() failed, %s", err)
	}
	if todoU.ID != updateTodo.ID || todoU.Note != updateTodo.Note || todoU.Title != updateTodo.Title || todoU.NoteDate.Equal(updateTodo.NoteDate) {
		t.Errorf("Test DbGetById() failed, gotten note and updated note is not equal")
	}
	rowsDeleted, err := postgres.Delete(strconv.Itoa(todoU.ID))
	if err != nil {
		t.Errorf("Test DbDelete() failed, %s", err)
	}
	if rowsDeleted != 1 {
		t.Errorf("Test DbDeleted() failed delete rows %d", rowsDeleted)
	}
}
