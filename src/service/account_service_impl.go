package service

import (
	"errors"
	"github.com/aasumitro/karlota/src/domain"
	"github.com/aasumitro/karlota/src/repository/mysql"
	"github.com/aasumitro/karlota/src/utils"
	"strconv"
	"time"
)

type accountServiceImpl struct {
	repo mysql.AccountRepository
}

func (acc accountServiceImpl) Register(user *domain.User) error {
	user.Password, _ = utils.Hash{}.Make(user.Password)

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

	// TODO: ADD TEST FOR THIS
	jwt := utils.JWT{}
	lifespan := time.Duration(86400) * time.Second
	tokenExpire := time.Now().Add(lifespan).Unix()
	token, err := jwt.Claim(user)
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

func AccountServiceImpl(repo mysql.AccountRepository) AccountService {
	return &accountServiceImpl{repo: repo}
}
