package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string `gorm:"primaryKey;unique;type:varchar(255)"`
	FirebaseUID string `gorm:"unique;type:varchar(255)"`
	AvatarURL   *string `gorm:"type:varchar(255)"`
	Name        string `gorm:"type:varchar(255)"`
}
