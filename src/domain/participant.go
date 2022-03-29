package domain

type Participant struct {
	ID             int64 `gorm:"column:id;primary_key" sql:"index" json:"id"`
	ConversationId int64 `gorm:"column:conversation_id" sql:"index" json:"conversation_id"`
	UserId         int64 `gorm:"column:user_id" sql:"index" json:"user_id"`
	TypingStatus   bool  `gorm:"column:typing_status" json:"typing_status"`
}

func (Participant) TableName() string {
	return "participants"
}
