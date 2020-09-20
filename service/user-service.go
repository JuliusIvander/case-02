package service

import (
	"github.com/JuliusIvander/case-02/entity"
)

type (
	// UserService interface
	UserService interface {
		GetUser(id string) (*entity.User, error)
		AddUser(*entity.User) error
	}
)
