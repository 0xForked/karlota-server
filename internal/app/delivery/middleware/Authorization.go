package middleware

import (
	utils2 "github.com/aasumitro/karlota/internal/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func Authorization(jwtUtils utils2.JSONWebToken) gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			utils2.NewHttpRespond(context, 401, "AUTHORIZATION_HEADER_REQUIRED")
			context.Abort()
			return
		}

		extractedToken := jwtUtils.ExtractFromHeader(authorizationHeader)
		token, err := jwtUtils.Verify(extractedToken)
		if err != nil {
			utils2.NewHttpRespond(context, http.StatusUnauthorized, err.Error())
			context.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			utils2.NewHttpRespond(context, http.StatusUnauthorized, "TOKEN_NOT_VALID")
			context.Abort()
			return
		}

		context.Set("payload", claims["payload"])
		context.Next()
	}
}
