package application

import (
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/domain"
	"github.com/atsushi-kitazawa/go_domain_driven_design_training/repository"
)

var (
	todoRepo repository.ITodoRepository
)

func Init(repo repository.ITodoRepository) {
	todoRepo = repo
}

func Register(todo domain.Todo) {
	todoRepo.Save(todo)
}

func Get(id int) domain.Todo {
	todo := todoRepo.Find(id)
	return todo
}
