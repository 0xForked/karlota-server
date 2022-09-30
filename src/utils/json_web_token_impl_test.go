package utils_test

import (
	"github.com/aasumitro/karlota/src/domain"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWT_Claim(t *testing.T) {
	type args struct {
		user *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			// TODO
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//j := &utils.JWT{}
			//got, err := j.Claim(tt.args.user)
			//if !tt.wantErr(t, err, fmt.Sprintf("Claim(%v)", tt.args.user)) {
			//	return
			//}
			//assert.Equalf(t, tt.want, got, "Claim(%v)", tt.args.user)
		})
	}
}

func TestJWT_ExtractFromHeader(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			// TODO
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//j := &utils.JWT{}
			//assert.Equalf(t, tt.want, j.ExtractFromHeader(tt.args.token), "ExtractFromHeader(%v)", tt.args.token)
		})
	}
}

func TestJWT_Verify(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *jwt.Token
		wantErr assert.ErrorAssertionFunc
	}{
		{
			// TODO
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//j := &utils.JWT{}
			//got, err := j.Verify(tt.args.token)
			//if !tt.wantErr(t, err, fmt.Sprintf("Verify(%v)", tt.args.token)) {
			//	return
			//}
			//assert.Equalf(t, tt.want, got, "Verify(%v)", tt.args.token)
		})
	}
}
