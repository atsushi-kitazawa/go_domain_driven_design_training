package domain

type Todo struct {
    Id int
    Name string
    Description string
}

func NewTodo(id int, name string, description string) *Todo {
    return &Todo{Id:id, Name:name, Description: description}
}

func (todo *Todo) ChangeName(name string) {
    todo.Name = name
}

func (todo *Todo) ChangeDescription(description string) {
    todo.Description = description
}
