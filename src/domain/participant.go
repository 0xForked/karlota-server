package domain

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Participant struct {
	ID             uint         `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	ConversationId uint         `gorm:"column:conversation_id;not null;index" json:"conversation_id"`
	UserId         uint         `gorm:"column:user_id;not null;index" json:"user_id"`
	TypingStatus   bool         `gorm:"column:typing_status;default:false;not null" json:"typing_status"`
	Conversation   Conversation `gorm:"foreignKey:conversation_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User           User         `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Participant) TableName() string {
	return "participants"
}

func (Participant) Migrate(db *gorm.DB) {
	if !db.Migrator().HasTable(Participant{}.TableName()) {
		if err := db.Migrator().AutoMigrate(Participant{}); err != nil {
			log.Panicln(fmt.Sprintf(
				"MIGRATE_ERROR(%s): %s",
				Participant{}.TableName(),
				err.Error(),
			))
		}
	}
}
