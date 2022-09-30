package conversation

import (
	"github.com/aasumitro/karlota/pkg/ws"
	"github.com/aasumitro/karlota/src/service"
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

func NewHandler(wsWrapper *ws.Melody, service service.AccountService, ) WsHandler {
	return &conversationHandler{
		wsWrapper: wsWrapper,
		service:   service,
	}
}

// list_chat, get_chat, start_chat
// add_user_to_chat (opt), remove_user_from_chat (opt), (chat) !if room
// set_online_status (chat)

// send_event [text, image, custom] (messages)
// send_typing_indicator (messages)
// mark_event_as_seen,

// list_chat: load user list chat
// get_chat: load selected chat thread (event: (message: text, image, custom), participants)
// start_chat: create new chat by with selected participants
// add_user_to_chat and remove_user_from_chat: add/remove user to chat for chat type room
// but, it is possible to add/remove user to chat for chat type private (just in case: video or voice)
// set_online_status: set online status for chat

// send_event: send message to chat
// send_typing_indicator: send typing indicator to chat
// mark_event_as_seen: mark event as seen

// start_chat
// DATA:
// 1. type (string)
// 2. participants (array of user_id (uint))
// FLOW:
// 1. create Conversation
// 2. create Participant

// list_chat
// DATA:
// 1. user_id (uint)
// FLOW:
// 1. Select Conversation where user_id is in Participant
// 2. Display Conversation by Type (room, private)
//		a. if private, display friends name
//		b. if room, display room name

// get_chat
// DATA:
// 1. user_id (uint)
// FLOW:
// 1. Select Message join Participant join User

// EVERY USER LOGIN OR CONNECT TO SOCKET SET IT ONLINE
//

// CLIENT
// LOGIN/REGISTER
// REDIRECT TO HOME
// IN HOME USER CONNECT TO SOCKET AND UPDATE HIS STATUS (ONLINE)
// USER SUBSCRIBE TO ONLINE USER LIST
// USER CLICK ONE OF THE LIST
// USER STREAM TO THE CONVERSATION
// USER SEND MESSAGE VIA CONVERSATION
// WITH SENDER_ID AND RECEIPANT_ID
// WHEN USER CLOSE THE APP
// SET USER AS OFFLINE
// TYPING STATUS WILL SENT DIRECTLY
// NOTIFICATION WILL USE FIREBASE (ON BACKGROUND)
// WEBSOCKET (ON FOREGROUND)
