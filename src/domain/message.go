package domain

import "database/sql/driver"

type MessageType string

//goland:noinspection ALL
const (
	Text   MessageType = "text"
	Image  MessageType = "image"
	Voice  MessageType = "voice"
	Custom MessageType = "custom"
)

func (t *MessageType) Scan(value interface{}) error {
	*t = MessageType(value.([]byte))
	return nil
}

func (t MessageType) Value() (driver.Value, error) {
	return string(t), nil
}

type Message struct {
	ID             int64       `gorm:"column:id;primary_key" sql:"index" json:"id"`
	ConversationId int64       `gorm:"column:conversation_id" sql:"index" json:"conversation_id"`
	ParticipantId  int64       `gorm:"column:participant_id" sql:"index" json:"user_id"`
	Type           MessageType `gorm:"column:type" sql:"type:ENUM('text','image','voice','custom');default:'text'" json:"type"`
	Body           string      `gorm:"column:body" json:"body"`
	SentAt         int64       `json:"sent_at"`
	SeenAt         int64       `json:"seen_at"`
	PulledAt       int64       `json:"pulled_at"`
}

func (Message) TableName() string {
	return "messages"
}
