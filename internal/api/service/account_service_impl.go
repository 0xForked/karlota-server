package service

import (
	"errors"
	"github.com/aasumitro/karlota/internal/api/domain"
	"github.com/aasumitro/karlota/internal/api/repository/mysql"
	"github.com/aasumitro/karlota/internal/utils"
	"strconv"
	"time"
)

type accountServiceImpl struct {
	repo mysql.AccountRepository
	jwt  utils.JSONWebToken
}

func (acc accountServiceImpl) Register(user *domain.User) error {
	user.Password = utils.Hash{}.Make(user.Password)

	return acc.repo.Store(user)
}

func (acc accountServiceImpl) Login(email string, password string) (interface{}, error) {
	user, err := acc.repo.Find(email)
	if err != nil {
		return nil, errors.New("EMAIL_NOT_FOUND")
	}

	verify := utils.Hash{}.Verify(password, user.Password)
	if !verify {
		return nil, errors.New("INVALID_PASSWORD")
	}

	lifespan := time.Duration(acc.jwt.GetExpirationHours()) * time.Hour
	tokenExpire := time.Now().Add(lifespan).Unix()
	token, err := acc.jwt.Claim(user)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"type":       "Bearer",
		"token":      token,
		"expired_in": strconv.FormatInt(tokenExpire, 10),
	}, nil
}

func (acc accountServiceImpl) Profile(email string) (*domain.User, error) {
	return acc.repo.Find(email)
}

func (acc accountServiceImpl) List() (*[]domain.User, error) {
	return acc.repo.All()
}

func (acc accountServiceImpl) Edit(user *domain.User) error {
	fcmToken := user.FCMToken
	newPassword := user.Password

	user, err := acc.repo.Find(user.Email)
	if err != nil {
		return errors.New("USER_NOT_FOUND")
	}

	if fcmToken != "" {
		user.FCMToken = fcmToken
	}

	if newPassword != "" {
		user.Password = utils.Hash{}.Make(newPassword)
	}

	if err := acc.repo.Update(user); err != nil {
		return errors.New("FAILED_UPDATE_DATA")
	}

	return nil
}

func AccountServiceImpl(repo mysql.AccountRepository, jwt utils.JSONWebToken) AccountService {
	return &accountServiceImpl{repo: repo, jwt: jwt}
}
