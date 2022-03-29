package mysql

import "github.com/aasumitro/karlota/src/domain"

type AccountRepository interface {
	Store(user *domain.User) error
	Find(email string) (*domain.User, error)
}
