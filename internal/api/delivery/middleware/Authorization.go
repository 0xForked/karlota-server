package middleware

import (
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func Authorization(jwtUtils utils.JSONWebToken) gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			utils.NewHttpRespond(context, 401, "AUTHORIZATION_HEADER_REQUIRED")
			context.Abort()
			return
		}

		extractedToken := jwtUtils.ExtractFromHeader(authorizationHeader)
		token, err := jwtUtils.Verify(extractedToken)
		if err != nil {
			utils.NewHttpRespond(context, http.StatusUnauthorized, err.Error())
			context.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			utils.NewHttpRespond(context, http.StatusUnauthorized, "TOKEN_NOT_VALID")
			context.Abort()
			return
		}

		context.Set("payload", claims["payload"])
		context.Next()
	}
}
