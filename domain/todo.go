package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Description string
	Completed   bool      `gorm:"default:false"`
	UserID      uuid.UUID `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (t *Todo) TableName() string {
	return "todos"
}

func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New()
	return nil
}
