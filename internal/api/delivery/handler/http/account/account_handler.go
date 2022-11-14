package account

import (
	"github.com/aasumitro/karlota/internal/api/delivery/middleware"
	"github.com/aasumitro/karlota/internal/api/service"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	router  *gin.Engine
	service service.AccountService
}

func NewHandler(router *gin.Engine, service service.AccountService, jwt utils.JSONWebToken) {
	handler := &accountHandler{router: router, service: service}

	v1 := handler.router.Group("/v1")
	{
		v1.POST("/register", handler.signUp)
		v1.POST("/login", handler.signIn)

		authorized := v1.Group("")
		authorized.Use(middleware.Authorization(jwt))
		{
			authorized.GET("/profile", handler.profile)
			authorized.GET("/users", handler.users)
			authorized.POST("/update/fcm", handler.updateFCMToken)
			authorized.POST("/update/password", handler.updatePassword)
			// authorized.GET("/logout", handler.signOut)
		}
	}
}
