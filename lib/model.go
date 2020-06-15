package todos

import "time"

type Config struct {
	ConnectionString string
	Port             int
}
type Todo struct {
	ID       int
	Title    string
	Note     string
	NoteDate time.Time
}
