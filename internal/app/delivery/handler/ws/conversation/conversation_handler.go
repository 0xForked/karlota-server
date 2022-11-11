package conversation

import (
	"github.com/aasumitro/karlota/internal/app/service"
	ws2 "github.com/aasumitro/karlota/internal/pkg/ws"
)

type WsHandler interface {
	OnConnected(session *ws2.Session)
	MessageHandler(s *ws2.Session, msg []byte)
	OnDisconnected(session *ws2.Session)
}

type conversationHandler struct {
	wsWrapper *ws2.Melody
	service   service.AccountService
}

func NewHandler(wsWrapper *ws2.Melody, service service.AccountService) WsHandler {
	return &conversationHandler{
		wsWrapper: wsWrapper,
		service:   service,
	}
}
