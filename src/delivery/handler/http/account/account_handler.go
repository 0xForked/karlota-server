package account

import (
	"github.com/aasumitro/karlota/src/delivery/middleware"
	"github.com/aasumitro/karlota/src/service"
	"github.com/aasumitro/karlota/src/utils"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	router  *gin.Engine
	service service.AccountService
}

func NewHandler(router *gin.Engine, service service.AccountService, jwt utils.JWT) {
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
			// authorized.GET("/logout", handler.signOut)
		}
	}
}
