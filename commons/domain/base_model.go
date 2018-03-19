package domain

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

type BaseModel struct {
	ID        string     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (user *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.New().String())
	return err
}
