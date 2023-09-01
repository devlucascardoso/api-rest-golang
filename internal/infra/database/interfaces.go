package database

import "github.com/devlucascardoso/api-rest-golang/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}