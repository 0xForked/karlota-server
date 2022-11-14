package listener

// FCMNotifyListener represent the data-struct for configuration
type FCMNotifyListener struct {
	//token       string
	//title       string
	//description string
	//fcmConn     *firebase.App
}

func (listener FCMNotifyListener) PushNotify() {}

func InitMessagingListener(
// token string,
// title string,
// description string,
// fcmConn *firebase.App,
) *FCMNotifyListener {
	return &FCMNotifyListener{
		//token: token,
		//title: title,
		//description: description,
		//fcmConn: fcmConn,
	}
}
