package ws

import (
	"fmt"
	"github.com/aasumitro/karlota/internal/api/delivery/handler/ws/conversation"
	"github.com/aasumitro/karlota/internal/api/repository/mysql"
	"github.com/aasumitro/karlota/internal/api/service"
	"github.com/aasumitro/karlota/internal/config"
	"github.com/aasumitro/karlota/internal/utils"
	"github.com/aasumitro/karlota/pkg/ws"
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
