package services

import (
	"fmt"

	"p3ld3v.dev/template/app/domain"
)

type UserService interface {
	GetUser(id int64) (*domain.User, error)
	CreateUser(name string) (*domain.User, error)
}

type UserServiceImp struct {
	db  DbStore
	log Logger
}

func NewUserService(store DbStore, log Logger) UserService {
	return &UserServiceImp{
		store,
		log,
	}
}

func (impl *UserServiceImp) GetUser(id int64) (*domain.User, error) {
	user, err := impl.db.GetUser(id)
	if err != nil {
		impl.log.Error(fmt.Sprintf("Error: %v\n", err))
		return nil, err
	}
	return &domain.User{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}

func (impl *UserServiceImp) CreateUser(name string) (*domain.User, error) {
	user, err := impl.db.CreateUser(name)
	if err != nil {
		impl.log.Error(fmt.Sprintf("Error: %v\n", err))
		return nil, err
	}
	return &domain.User{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}
