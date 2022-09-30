package conversation

import (
	"fmt"
	"github.com/aasumitro/karlota/internal/pkg/ws"
)

func (handler *conversationHandler) OnDisconnected(session *ws.Session) {
	fmt.Println("Disconnected Set Offline", session.Keys)
}
