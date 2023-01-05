// Package user package user
package user

import (
	"tpc/internal/entity"
	"tpc/internal/usecase"
)

type UserUseCase struct{}

func NewUserUseCase() usecase.User {
	return &UserUseCase{}
}

func (uc *UserUseCase) GetUserByUserName(userName string) (entity.User, error) {
	return entity.User{}, nil
}

func (uc *UserUseCase) UpdateUserByUsername(userName string, user entity.User) error {
	return nil
}

func (uc *UserUseCase) InsertUser(record *entity.User) error {
	return nil
}
