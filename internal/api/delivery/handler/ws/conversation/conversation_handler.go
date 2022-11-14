package conversation

import (
	"github.com/aasumitro/karlota/internal/api/service"
	"github.com/aasumitro/karlota/pkg/ws"
)

type WsHandler interface {
	OnConnected(session *ws.Session)
	MessageHandler(s *ws.Session, msg []byte)
	OnDisconnected(session *ws.Session)
}

type conversationHandler struct {
	wsWrapper *ws.Melody
	service   service.AccountService
}

func NewHandler(wsWrapper *ws.Melody, service service.AccountService) WsHandler {
	return &conversationHandler{
		wsWrapper: wsWrapper,
		service:   service,
	}
}
