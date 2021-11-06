package repository

import (
	"database/sql"
	_ "fmt"

	"github.com/atsushi-kitazawa/go_domain_driven_design_training/domain"
)

type ITodoRepository interface {
	Find(id int)
	Save(todo domain.Todo)
}

type TodoRepository struct {
}

type InMemoryTodoRepository struct {
    Stores []domain.Todo
}

func (t TodoRepository) Find(id int) domain.Todo {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=15432 user=postgres password=mysecretpassword dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, name, description from todo")
	if err != nil {
		panic(err)
	}

	var todo domain.Todo
	for rows.Next() {
		rows.Scan(&todo.Id, &todo.Name, &todo.Description)
	}

	return todo
}

func (t TodoRepository) Save(todo domain.Todo) {
    db, err := sql.Open("postgres", "host=127.0.0.1 port=15432 user=postgres password=mysecretpassword dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

    stmt, err := db.Prepare("insert into todo values ($1, $2, $3)")
    if err != nil {
	panic(err)
    }

    _, err = stmt.Exec(todo.Id, todo.Name, todo.Description)
    if err != nil {
	panic(err)
    }
}
