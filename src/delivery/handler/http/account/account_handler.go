package account

import (
	"github.com/aasumitro/karlota/src/delivery/middleware"
	"github.com/aasumitro/karlota/src/service"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	router  *gin.Engine
	service service.AccountService
}

func NewHandler(router *gin.Engine, service service.AccountService) {
	handler := &accountHandler{router: router, service: service}

	v1 := handler.router.Group("/v1")
	{
		v1.POST("/register", handler.register)
		v1.POST("/login", handler.login)

		authorized := v1.Group("")
		authorized.Use(middleware.Authorization)
		{
			authorized.GET("/profile", handler.profile).Use()
		}
	}
}
