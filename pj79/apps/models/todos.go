package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID         int
	Content    string
	UserID     int
	created_at time.Time
}

func (u *User) CreateTodos(content string) (err error) {
	cmd := fmt.Sprintf(`INSERT INTO todos (
		content,
		user_id,
		created_at) VALUES (?, ?, ?)`)

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTod(id int) (todo Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE ID = ?`
}