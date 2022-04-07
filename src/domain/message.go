package domain

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"log"
)

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
	ID             uint         `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	ConversationId uint         `gorm:"column:conversation_id;not null;index" json:"conversation_id"`
	ParticipantId  uint         `gorm:"column:participant_id;not null;index" json:"participant_id"`
	Type           MessageType  `gorm:"column:type;type:ENUM('text','image','voice','custom');default:'text'" json:"type"`
	Body           string       `gorm:"column:body" json:"body"`
	SentAt         sql.NullTime `gorm:"column:sent_at;" json:"sent_at"`
	SeenAt         sql.NullTime `gorm:"column:seen_at;" json:"seen_at"`
	PulledAt       sql.NullTime `gorm:"column:pulled_at;" json:"pulled_at"`
	Conversation   Conversation `gorm:"foreignKey:conversation_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Participant    Participant  `gorm:"foreignKey:participant_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Message) TableName() string {
	return "messages"
}

func (Message) Migrate(db *gorm.DB) {
	if !db.Migrator().HasTable(Message{}.TableName()) {
		if err := db.Migrator().AutoMigrate(Message{}); err != nil {
			log.Panicln(fmt.Sprintf(
				"MIGRATE_ERROR(%s): %s",
				Message{}.TableName(),
				err.Error(),
			))
		}
	}
}
