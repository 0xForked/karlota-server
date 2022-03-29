package middleware

import (
	"github.com/aasumitro/karlota/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorization(context *gin.Context) {
	authorizationHeader := context.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		utils.NewHttpRespond(context, 401, "AUTHORIZATION_HEADER_REQUIRED")
		context.Abort()
		return
	}

	jwtUtils := utils.JWT{}
	extractedToken := jwtUtils.ExtractFromHeader(authorizationHeader)
	token, err := jwtUtils.Verify(extractedToken)
	if err != nil {
		utils.NewHttpRespond(context, 401, err.Error())
		context.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		utils.NewHttpRespond(context, 401, "TOKEN_NOT_VALID")
		context.Abort()
		return
	}

	context.Set("payload", claims["payload"])
	context.Next()
}
