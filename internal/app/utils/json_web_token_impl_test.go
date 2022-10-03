package utils_test

import (
	"errors"
	"github.com/aasumitro/karlota/internal/app/domain"
	"github.com/aasumitro/karlota/internal/app/utils"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWT_Claim(t *testing.T) {
	jwtUtil := utils.NewJWTUtil("secret", "KARLOTA SERVER", 24)
	user := domain.User{ID: 1, Name: "lorem", Email: "lorem@ipsum.id"}
	token, err := jwtUtil.Claim(&user)
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestJWT_Verify(t *testing.T) {
	jwtUtil := utils.NewJWTUtil("secret", "KARLOTA SERVER", 24)
	extractedToken := jwtUtil.ExtractFromHeader("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQ2MTQwOTYsImlhdCI6MTY2NDUyNzY5NiwiaXNzIjoiS0FSTE9UQSBTRVJWRVIiLCJwYXlsb2FkIjp7ImlkIjoyLCJuYW1lIjoiYWFzdW1pdHJvIiwiZW1haWwiOiJoZWxsb0BhYXN1bWl0cm8uaWQiLCJmY21fdG9rZW4iOiIiLCJpc19vbmxpbmUiOmZhbHNlfX0.0uw7TyV1cXtHQ7Lj2e6vMq1uTR_wbYjBih9juAyRYSU")
	token, err := jwtUtil.Verify(extractedToken)
	assert.Equal(t, "Token is expired", err.Error())
	claims, ok := token.Claims.(jwt.MapClaims)
	assert.Equal(t, true, ok)
	assert.Equal(t, "KARLOTA SERVER", claims["iss"])
}

func TestJWT_Verify_InvalidSigningMethodHS256(t *testing.T) {
	jwtUtil := utils.NewJWTUtil("secret", "KARLOTA SERVER", 24)
	token := jwtUtil.ExtractFromHeader("Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQ2MTYyNzUsImlhdCI6MTY2NDUyOTg3NSwiaXNzIjoiS0FSTE9UQSBTRVJWRVIiLCJwYXlsb2FkIjp7ImlkIjoyLCJuYW1lIjoiYWFzdW1pdHJvIiwiZW1haWwiOiJoZWxsb0BhYXN1bWl0cm8uaWQiLCJmY21fdG9rZW4iOiIxMjMtYWFzLTQ1NiIsImlzX29ubGluZSI6ZmFsc2V9fQ.HIm0UPfC7BV-hmXgrXijXUNLT4qxjeuo39V3_BTgqekyv2W2PIGi9tBhYMtuILYcExaI3Fl53NbeEHQDnAWalA")
	_, err := jwtUtil.Verify(token)
	assert.Error(t, errors.New("signature is invalid"), err)
}

func TestJWT_Verify_InvalidSignature(t *testing.T) {
	jwtUtil := utils.NewJWTUtil("secret", "KARLOTA SERVER", 24)
	token := jwtUtil.ExtractFromHeader("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQ2MTQwOTYsImlhdCI6MTY2NDUyNzY5NiwiaXNzIjoiS0FSTE9UQSBTRVJWRVIiLCJwYXlsb2FkIjp7ImlkIjoyLCJuYW1lIjoiYWFzdW1pdHJvIiwiZW1haWwiOiJoZWxsb0BhYXN1bWl0cm8uaWQiLCJmY21fdG9rZW4iOiIiLCJpc19vbmxpbmUiOmZhbHNlfX0.0uw7TyV1cXtHQ7Lj2e6vMq1uTR_wbYjBih9juAyRYSUs")
	_, err := jwtUtil.Verify(token)
	assert.Error(t, errors.New("signature is invalid"), err)
}

func TestJWT_ExtractFromHeader_ShouldSuccess(t *testing.T) {
	jwtUtil := utils.NewJWTUtil("secret", "KARLOTA SERVER", 24)
	data := jwtUtil.ExtractFromHeader("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQ2MTQwOTYsImlhdCI6MTY2NDUyNzY5NiwiaXNzIjoiS0FSTE9UQSBTRVJWRVIiLCJwYXlsb2FkIjp7ImlkIjoyLCJuYW1lIjoiYWFzdW1pdHJvIiwiZW1haWwiOiJoZWxsb0BhYXN1bWl0cm8uaWQiLCJmY21fdG9rZW4iOiIiLCJpc19vbmxpbmUiOmZhbHNlfX0.0uw7TyV1cXtHQ7Lj2e6vMq1uTR_wbYjBih9juAyRYSU")
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQ2MTQwOTYsImlhdCI6MTY2NDUyNzY5NiwiaXNzIjoiS0FSTE9UQSBTRVJWRVIiLCJwYXlsb2FkIjp7ImlkIjoyLCJuYW1lIjoiYWFzdW1pdHJvIiwiZW1haWwiOiJoZWxsb0BhYXN1bWl0cm8uaWQiLCJmY21fdG9rZW4iOiIiLCJpc19vbmxpbmUiOmZhbHNlfX0.0uw7TyV1cXtHQ7Lj2e6vMq1uTR_wbYjBih9juAyRYSU", data)
}

func TestJWT_ExtractFromHeader_ShouldError(t *testing.T) {
	jwtUtil := utils.NewJWTUtil("secret", "KARLOTA SERVER", 24)
	data := jwtUtil.ExtractFromHeader("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQ2MTQwOTYsImlhdCI6MTY2NDUyNzY5NiwiaXNzIjoiS0FSTE9UQSBTRVJWRVIiLCJwYXlsb2FkIjp7ImlkIjoyLCJuYW1lIjoiYWFzdW1pdHJvIiwiZW1haWwiOiJoZWxsb0BhYXN1bWl0cm8uaWQiLCJmY21fdG9rZW4iOiIiLCJpc19vbmxpbmUiOmZhbHNlfX0.0uw7TyV1cXtHQ7Lj2e6vMq1uTR_wbYjBih9juAyRYSU")
	assert.Equal(t, "INVALID_TOKEN_FORMAT", data)
}

func TestJwtData_GetExpirationHours(t *testing.T) {
	jwtUtil := utils.NewJWTUtil("123", "test", 1)
	expHrs := jwtUtil.GetExpirationHours()
	assert.Equal(t, 1, expHrs)
}
