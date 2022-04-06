package utils

import (
	"fmt"
	"github.com/aasumitro/karlota/src/domain"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type JWT struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type MyJWTClaim struct {
	jwt.StandardClaims
	Payload interface{} `json:"payload"`
}

func (j *JWT) Claim(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.Issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(j.ExpirationHours) * time.Hour).Unix(),
		},
		Payload: user,
	})

	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWT) Verify(signedToken string) (*jwt.Token, error) {
	return jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("INVALID_TOKEN: %s", token.Header["alg"])
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("INVALID_SIGNING_METHOD: %s", method.Alg())
		}

		return []byte(j.SecretKey), nil
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
