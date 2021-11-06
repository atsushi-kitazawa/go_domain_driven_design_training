package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/go_domain_driven_design_training/domain"
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/infra"
)

func main() {
    fmt.Println("main")
    test2()
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
