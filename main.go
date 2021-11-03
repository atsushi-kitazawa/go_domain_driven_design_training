package main

import (
	"fmt"

	"github.com/atsushi-kitazawa/go_domain_driven_design_training/domain"
)

func main() {
    fmt.Println("main")

    user, err := domain.NewUser("hoge", "090-0000-0000")
    if err != nil {
	panic(err)
    }
    fmt.Println(user)
}
