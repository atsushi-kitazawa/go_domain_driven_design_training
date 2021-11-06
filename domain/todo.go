package domain

type Todo struct {
    id int
    name string
    description string
}

func NewTodo(id int, name string, description string) *Todo {
    return &Todo{id:id, name:name, description: description}
}

func (todo *Todo) ChangeName(name string) {
    todo.name = name
}

func (todo *Todo) ChangeDescription(description string) {
    todo.description = description
}
