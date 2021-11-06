package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/go_domain_driven_design_training/domain"
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/infra"
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/repository"
)

func main() {
	fmt.Println("main")
	//test2()
	//test3()
	test4()
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

func test4() {
	repo := repository.TodoRepository{}
	todo := repo.Find(1)
	fmt.Println(todo)

	todo = *domain.NewTodo(2, "todoB", "this todo is B")
	repo.Save(todo)

	repo = repository.TodoRepository{}
	todo = repo.Find(2)
	fmt.Println(todo)

}
