package ws

import (
	"github.com/aasumitro/karlota/config"
	"github.com/aasumitro/karlota/pkg/ws"
	"github.com/gin-gonic/gin"
)

type wsHandler struct{}

func NewWsHandler(config *config.Config, router *gin.Engine) {
	//handler := &wsHandler{}
	m := ws.New()

	router.GET("/conversation/:id/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *ws.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *ws.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
}
