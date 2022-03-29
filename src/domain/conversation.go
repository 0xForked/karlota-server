package domain

import (
	"database/sql/driver"
)

type ConversationType string

//goland:noinspection ALL
const (
	Private ConversationType = "private"
	Group   ConversationType = "group"
)

func (t *ConversationType) Scan(value interface{}) error {
	*t = ConversationType(value.([]byte))
	return nil
}

func (t ConversationType) Value() (driver.Value, error) {
	return string(t), nil
}

type Conversation struct {
	ID   int64            `gorm:"column:id;primary_key" sql:"index" json:"id"`
	Type ConversationType `gorm:"column:type" sql:"type:ENUM('private', 'group');default:'private'" json:"type"`
	// name & avatar will be NULL when Type is Private
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	// DATE with Soft-delete and archived
	CreatedAt  int64 `json:"created_at"`
	UpdatedAt  int64 `json:"updated_at"`
	DeletedAt  int64 `json:"deleted_at"`
	ArchivedAt int64 `json:"archived_at"`
	// RELATIONSHIPS
	// 1:n Participants, 1:n Messages
}

func (Conversation) TableName() string {
	return "conversations"
}
