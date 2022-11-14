package service

import (
	"github.com/aasumitro/karlota/internal/api/domain"
)

type AccountService interface {
	Edit(user *domain.User) error
	Register(user *domain.User) error
	Login(email string, password string) (interface{}, error)
	Profile(email string) (*domain.User, error)
	List() (*[]domain.User, error)
}
