package db

import (
	"fmt"
)

type UserRepository interface {
	Create() error
} 

type UserRepositoryImpl struct {
	
}

func NewUserRepository () UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create() error {
	fmt.Println("Creating user")
	return nil
}
