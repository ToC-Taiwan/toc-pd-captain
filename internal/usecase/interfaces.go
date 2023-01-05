// Package usecase package usecase
package usecase

import "tpc/internal/entity"

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

// User -.
type User interface {
	GetUserByUserName(userName string) (entity.User, error)
	UpdateUserByUsername(userName string, user entity.User) error
	InsertUser(record *entity.User) error
}
