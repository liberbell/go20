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
		created_at) VALUES ($1, $2, $3)`)

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE ID = ?`

	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.created_at)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.created_at)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.created_at)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `UPDATE todos SET content = ?, user_id = ? WHERE id = ?`
	_, err := Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `DELETE FROM todos WHERE id = ?`
	_, err := Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
