package repository

import (
	"database/sql"
	_ "fmt"

	"github.com/atsushi-kitazawa/go_domain_driven_design_training/domain"
)

type ITodoRepository interface {
	Find(id int) domain.Todo
	Save(todo domain.Todo)
	Delete(id int)
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

func (t TodoRepository) Delete(id int) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=15432 user=postgres password=mysecretpassword dbname=testdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from todo where id = $1")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}

func NewInmemoryTodoRepository() *InMemoryTodoRepository {
    ret := &InMemoryTodoRepository{}
    ret.Stores = make([]domain.Todo, 10)
    return ret
}

func (t *InMemoryTodoRepository) Find(id int) domain.Todo {
    for _, v := range t.Stores {
	if v.Id == id {
	    return v
	}
    }
    return domain.Todo{}
}

func (t *InMemoryTodoRepository) Save(todo domain.Todo) {
    t.Stores = append(t.Stores, todo)
}

func (t *InMemoryTodoRepository) Delete(id int) {
    for i, v := range t.Stores {
	if v.Id == id {
	    t.Stores[i] = t.Stores[len(t.Stores)-1]
	}
    }
}
