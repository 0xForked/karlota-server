package conversation

// list_chat, get_chat, start_chat
// add_user_to_chat (opt), remove_user_from_chat (opt), (chat) !if room
// set_online_status (chat)

// send_event [text, image, custom] (messages)
// send_typing_indicator (messages)
// mark_event_as_seen,

// list_chat: load user list chat
// get_chat: load selected chat thread (events: (message: text, image, custom), participants)
// start_chat: create new chat by with selected participants
// add_user_to_chat and remove_user_from_chat: add/remove user to chat for chat type room
// but, it is possible to add/remove user to chat for chat type private (just in case: video or voice)
// set_online_status: set online status for chat

// send_event: send message to chat
// send_typing_indicator: send typing indicator to chat
// mark_event_as_seen: mark event as seen
