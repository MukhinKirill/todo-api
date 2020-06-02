package db

import (
	"database/sql"

	model "../model"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func ConnectDb(connectionString string) (*Postgres, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{db}, nil
}
func (p *Postgres) Close() {
	p.DB.Close()
}

func (p *Postgres) DbInit() (int, error) {
	query := `
		CREATE SEQUENCE IF NOT EXISTS todo_id START 1
	`
	_, err := p.DB.Exec(query)
	if err != nil {
		return -1, err
	}
	query = `
	CREATE TABLE IF NOT EXISTS todo(
		ID int PRIMARY KEY,
		TITLE TEXT NOT NULL,
		NOTE TEXT,
		NOTE_DATE TIMESTAMP WITH TIME ZONE)
		`
	_, err = p.DB.Exec(query)
	if err != nil {
		return -1, err
	}
	return 0, nil
}

func (p *Postgres) Insert(todo *model.Todo) (int, error) {
	query := `
		INSERT INTO todo (id, title, note, note_date)
		VALUES (nextval('todo_id'), $1, $2, $3)
		RETURNING id;
	`

	rows, err := p.DB.Query(query, todo.Title, todo.Note, todo.NoteDate)
	if err != nil {
		return -1, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}

	return id, nil
}

func (p *Postgres) Update(todo *model.Todo, idStr string) (int, error) {
	query := `
		UPDATE todo SET title = $2, note= $3, note_date = $4
		WHERE id = $1
		RETURNING id;
	`

	rows, err := p.DB.Query(query, idStr, todo.Title, todo.Note, todo.NoteDate)
	if err != nil {
		return -1, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}

	return id, nil
}

func (p *Postgres) Delete(id string) error {
	query := `
		DELETE FROM todo
		WHERE id = $1;
	`

	if _, err := p.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetAll() ([]model.Todo, error) {
	query := `
		SELECT *
		FROM todo
		ORDER BY id;
	`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var todoList []model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Note, &t.NoteDate); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}

	return todoList, nil
}

func (p *Postgres) GetOne(id string) (model.Todo, error) {
	query := `
		SELECT *
		FROM todo
		WHERE id=$1;
	`

	var todo model.Todo
	rows, err := p.DB.Query(query, id)
	if err != nil {
		return todo, err
	}

	var todoList []model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Note, &t.NoteDate); err != nil {
			return todo, err
		}
		todoList = append(todoList, t)
	}

	return todoList[0], nil
}
