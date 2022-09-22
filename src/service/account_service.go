package service

import (
	"github.com/aasumitro/karlota/src/domain"
)

type AccountService interface {
	Register(user *domain.User) error
	Login(email string, password string) (interface{}, error)
	Profile(email string) (*domain.User, error)
	List() (*[]domain.User, error)
}
