package utils

import (
	"fmt"
	"github.com/aasumitro/karlota/src/domain"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JWT struct{}

type myClaim struct {
	jwt.StandardClaims
	Payload interface{} `json:"payload"`
}

var (
	secret = "secret$%^!@12345://@()"
	exp    = time.Duration(86400) * time.Second
	iss    = "KARLOTA"
)

func (j *JWT) Claim(user *domain.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    iss,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
		Payload: user,
	})

	return token.SignedString([]byte(secret))
}

func (j *JWT) Verify(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("INVALID_TOKEN: %s", token.Header["alg"])
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("INVALID_SIGNING_METHOD: %s", method.Alg())
		}

		return []byte(secret), nil
	})
}

// ExtractFromHeader SendFrom Middleware c.Request.Header.Get("Authorization")
func (j *JWT) ExtractFromHeader(token string) string {
	tokenHeadName := "Bearer"
	parts := strings.SplitN(token, " ", 2)

	if (len(parts) == 2) && (parts[0] == tokenHeadName) {
		return strings.Split(token, " ")[1]
	}

	return "INVALID_TOKEN_FORMAT"
}
