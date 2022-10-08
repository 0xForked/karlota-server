package cmd

import (
	"github.com/aasumitro/karlota/docs"
	httpDelivery "github.com/aasumitro/karlota/internal/app/delivery/handler/http"
	rtcDelivery "github.com/aasumitro/karlota/internal/app/delivery/handler/rtc"
	wsDelivery "github.com/aasumitro/karlota/internal/app/delivery/handler/ws"
	"github.com/aasumitro/karlota/internal/config"
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
	docs.SwaggerInfo.Title = appConfig.GetAppName()
	docs.SwaggerInfo.Description = appConfig.GetAppDesc()
	docs.SwaggerInfo.Version = appConfig.GetAppVersion()
	docs.SwaggerInfo.Host = appConfig.GetAppUrl()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func StartServer() {
	httpDelivery.NewHttpHandler(appConfig, ginEngine)

	wsDelivery.NewWsHandler(appConfig, ginEngine)

	rtcDelivery.NewWebRTCHandler(appConfig)

	log.Fatal(ginEngine.Run(appConfig.GetAppUrl()))
}
