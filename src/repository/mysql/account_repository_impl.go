package mysql

import (
	"github.com/aasumitro/karlota/src/domain"
	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	db *gorm.DB
}

func init() {
	// TODO: implement
	// User Table Migration
	// and if table is new created
	// run data seeder
}

func (acc accountRepositoryImpl) Store(user *domain.User) error {
	if err := acc.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (acc accountRepositoryImpl) Find(email string) (*domain.User, error) {
	var user domain.User

	if err := acc.db.First(&user, `email = ?`, email).Error; err != nil {
		return &user, err
	}

	return &user, nil
}

func AccountRepositoryImpl(db *gorm.DB) AccountRepository {
	return &accountRepositoryImpl{db: db}
}
