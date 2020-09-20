package repository

import (
	"github.com/JuliusIvander/case-02/entity"
)

type (
	// UserRepositry interface
	UserRepositry interface {
		GetUserByID(id string) (*entity.User, error)
		AddUser(*entity.User) error
	}
)
