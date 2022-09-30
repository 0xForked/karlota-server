package conversation

import (
	"fmt"
	"github.com/aasumitro/karlota/pkg/ws"
)

func (handler *conversationHandler) OnDisconnected(session *ws.Session) {
	fmt.Println("Disconnected Set Offline", session.Keys)
}
