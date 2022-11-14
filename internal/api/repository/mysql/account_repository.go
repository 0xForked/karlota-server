package mysql

import (
	"github.com/aasumitro/karlota/internal/api/domain"
)

type AccountRepository interface {
	Update(user *domain.User) error
	Store(user *domain.User) error
	Find(email string) (*domain.User, error)
	All() (*[]domain.User, error)
}
