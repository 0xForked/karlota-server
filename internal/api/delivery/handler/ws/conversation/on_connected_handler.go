package conversation

import (
	"fmt"
	"github.com/aasumitro/karlota/pkg/ws"
)

// SEND BACK EXCHANGE CODE
// FROM FRONT-END CALL AUTH
// SEND EXCHANGE CODE + EMAIL

func (handler *conversationHandler) OnConnected(session *ws.Session) {
	fmt.Println(session.Request.Header)
	fmt.Println("Connected Set Online")
}
