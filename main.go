package main

import (
	"flag"
	"fmt"

	"github.com/atsushi-kitazawa/go_domain_driven_design_training/application"
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/domain"
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/infra"
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/repository"
)

func main() {
    // parse args
    id := flag.Int("id", 0, "id")
    name := flag.String("name", "todo", "todo name")
    description := flag.String("desc", "todo description", "todo description")
    flag.Parse()

    // init application service
    application.Init(repository.TodoRepository{})

    // register todo
    todo := domain.NewTodo(*id, *name, *description)
    application.Register(*todo)

    // get todo
    getTodo := application.Get(*id)
    fmt.Println(getTodo)
}

func test1() {
	user, err := domain.NewUser("hoge", "090-0000-0000")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	user = user.ChangeUserName("hoge change")
	fmt.Println(user)

	user = user.ChangeUserName("h")
	fmt.Println(user)
}

func test2() {
	infra.Select()
}

func test3() {
	todo := domain.NewTodo(1, "'aaa", "aaa description")
	fmt.Println(todo)
}
