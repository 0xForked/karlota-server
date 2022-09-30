package conversation

import (
	"fmt"
	"github.com/aasumitro/karlota/internal/pkg/ws"
)

func (handler *conversationHandler) MessageHandler(s *ws.Session, msg []byte) {
	if err := handler.wsWrapper.BroadcastFilter(msg, func(q *ws.Session) bool {
		return q.Request.URL.Path == s.Request.URL.Path
	}); err != nil {
		fmt.Println(err)
	}
}
