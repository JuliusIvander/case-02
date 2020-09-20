package service

import (
	"github.com/JuliusIvander/case-02/entity"
	"github.com/JuliusIvander/case-02/repository"
	"github.com/segmentio/ksuid"
)

type (
	userServiceImpl struct {
		repo repository.UserRepositry
	}
)

// NewUserService function
func NewUserService(repo repository.UserRepositry) UserService {
	return &userServiceImpl{repo}
}

// GetUser method
func (m *userServiceImpl) GetUser(id string) (*entity.User, error) {
	user, err := m.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// AddUser method
func (m *userServiceImpl) AddUser(user *entity.User) error {
	user.ID = ksuid.New().String()
	return m.repo.AddUser(user)
}
