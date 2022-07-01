package domain

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
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

// Conversation
// name & avatar will be NULL when Type is Private
// DATE with Soft-delete
type Conversation struct {
	ID           uint             `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	Type         ConversationType `gorm:"column:type;type:ENUM('private','group');default:'private'" json:"type"`
	Name         sql.NullString   `gorm:"column:name" json:"name"`
	Avatar       sql.NullString   `gorm:"column:avatar" json:"avatar"`
	CreatedAt    time.Time        `gorm:"column:created_at;index;autoCreateTime" json:"created_at"`
	UpdatedAt    sql.NullTime     `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt   `gorm:"column:deleted_at;index;"`
	Participants []Participant    `gorm:"foreignKey:conversation_id"`
	Messages     []Message        `gorm:"foreignKey:conversation_id"`
}

func (Conversation) TableName() string {
	return "conversations"
}

func (Conversation) Migrate(db *gorm.DB) {
	if !db.Migrator().HasTable(Conversation{}.TableName()) {
		if err := db.Migrator().AutoMigrate(Conversation{}); err != nil {
			log.Panicln(fmt.Sprintf(
				"MIGRATE_ERROR(%s): %s",
				Conversation{}.TableName(),
				err.Error(),
			))
		}
	}
}
