package ws

import (
	"fmt"
	"github.com/aasumitro/karlota/internal/app/delivery/handler/ws/conversation"
	"github.com/aasumitro/karlota/internal/app/repository/mysql"
	"github.com/aasumitro/karlota/internal/app/service"
	"github.com/aasumitro/karlota/internal/app/utils"
	"github.com/aasumitro/karlota/internal/config"
	"github.com/aasumitro/karlota/internal/pkg/ws"
	"github.com/gin-gonic/gin"
)

func NewWsHandler(config *config.Config, router *gin.Engine) {
	m := ws.New()
	accountRepository := mysql.AccountRepositoryImpl(
		config.GetDbConn())
	jwtUtils := utils.NewJWTUtil(
		config.GetJWTSecretKey(),
		config.GetAppName(),
		config.GetJWTLifespan(),
	)
	accountService := service.AccountServiceImpl(
		accountRepository, jwtUtils)
	conversationHandler := conversation.NewHandler(m, accountService)

	m.HandleConnect(conversationHandler.OnConnected)

	router.GET("/conversation", func(c *gin.Context) {
		if err := m.HandleRequest(c.Writer, c.Request); err != nil {
			fmt.Println(err)
		}
	})

	m.HandleMessage(conversationHandler.MessageHandler)

	m.HandleDisconnect(conversationHandler.OnDisconnected)
}
