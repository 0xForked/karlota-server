package cmd

import (
	"github.com/aasumitro/karlota/config"
	"github.com/aasumitro/karlota/docs"
	httpDelivery "github.com/aasumitro/karlota/src/delivery/handler/http"
	"github.com/gin-gonic/gin"
	"log"
)

var appConfig *config.Config

var ginEngine *gin.Engine

func init() {
	appConfig = &config.Config{}
	appConfig.InitDbConn()

	if !appConfig.GetAppDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine = gin.Default()

	docs.SwaggerInfo.BasePath = ginEngine.BasePath()
}

// StartServer
// @title KARLOTA - Instant Messaging Service Example
// @version 1.0
// @description REST, WebRTC, WebSocket.
// @termsOfService http://swagger.io/terms/
// @contact.name @aasumitro
// @contact.url https://aasumitro.id/
// @contact.email hello@aasumitro.id
// @BasePath /api/v1
// @license.name  MIT
// @license.url   https://github.com/aasumitro/karlota/blob/master/LICENSE
func StartServer() {
	httpDelivery.NewHttpHandler(appConfig, ginEngine)

	log.Fatal(ginEngine.Run(appConfig.GetAppUrl()))
}
