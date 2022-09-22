package event

type PushNotificationEvent struct{}

func (event PushNotificationEvent) Handler() {}

func InitPushNotificationEvent() *PushNotificationEvent {
	return &PushNotificationEvent{}
}
