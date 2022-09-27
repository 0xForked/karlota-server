package ws

import (
	"fmt"
	"github.com/aasumitro/karlota/config"
	"github.com/aasumitro/karlota/pkg/ws"
	"github.com/gin-gonic/gin"
)

type wsHandler struct{}

func NewWsHandler(config *config.Config, router *gin.Engine) {
	m := ws.New()

	router.GET("/conversation", func(c *gin.Context) {
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err)
		}
	})

	m.HandleMessage(func(s *ws.Session, msg []byte) {
		if err := m.BroadcastFilter(msg, func(q *ws.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		}); err != nil {
			fmt.Println(err)
		}
	})
}
