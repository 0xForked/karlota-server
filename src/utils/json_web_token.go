package utils

import (
	"github.com/aasumitro/karlota/src/domain"
	"github.com/golang-jwt/jwt"
)

type JSONWebToken interface {
	Claim(user *domain.User) (string, error)
	Verify(signedToken string) (*jwt.Token, error)
	ExtractFromHeader(token string) string
	GetExpirationHours() int
}
