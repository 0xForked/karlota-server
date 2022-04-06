package http

import (
	"fmt"
	"github.com/aasumitro/karlota/config"
	"github.com/aasumitro/karlota/src/delivery/handler/http/account"
	"github.com/aasumitro/karlota/src/repository/mysql"
	"github.com/aasumitro/karlota/src/service"
	"github.com/aasumitro/karlota/src/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// main route handler
type httpHandler struct{}

func NewHttpHandler(config *config.Config, router *gin.Engine) {
	handler := &httpHandler{}
	router.NoMethod(handler.noMethod)
	router.NoRoute(handler.notFound)
	router.GET("/", handler.home)
	router.GET("/ping", handler.ping)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Account Handler (Auth & Profile)
	accountRepository := mysql.AccountRepositoryImpl(config.GetDbConn())
	jwtUtils := utils.JWT{
		SecretKey:       config.GetJWTSecretKey(),
		ExpirationHours: config.GetJWTLifespan(),
		Issuer:          config.GetAppName(),
	}
	accountService := service.AccountServiceImpl(accountRepository, jwtUtils)
	account.NewHandler(router, accountService, jwtUtils)
}

func (handler httpHandler) home(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusOK, map[string]interface{}{
		"01_title":       "Karlota",
		"02_description": " Instant Messaging Service Example",
		"03_api_spec":    fmt.Sprintf("http://%s/docs/index.html", context.Request.Host),
		"04_perquisites": map[string]interface{}{
			"01_language":  "https://github.com/golang/go",
			"02_framework": "https://github.com/gin-gonic/gin",
			"03_library": map[string]string{
				"01_swagger": "https://github.com/swaggo/swag",
			},
		},
	})
}

func (handler httpHandler) ping(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusOK, "PONG")
}

func (handler httpHandler) notFound(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusNotFound, "HTTP_ROUTE_NOT_FOUND")
}

func (handler httpHandler) noMethod(context *gin.Context) {
	utils.NewHttpRespond(context, http.StatusNotFound, "HTTP_METHOD_NOT_FOUND")
}
