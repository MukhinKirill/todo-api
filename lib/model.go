package todos

import "time"

type Config struct {
	ConnectionString string
	Port             int
}
type Todo struct {
	ID       int
	Title    string    `valid:"type(string),required,length(1|255)"`
	Note     string    `valid:"type(string),required,length(0|2000)`
	NoteDate time.Time `valid:"type(time.Time),required"`
}
