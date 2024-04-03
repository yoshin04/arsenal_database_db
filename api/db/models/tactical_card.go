package models

import (
	"gorm.io/gorm"
)

type TacticalCard struct {
	gorm.Model
	ID             string `gorm:"primaryKey;unique;type:varchar(255)"`
	ImageUrl       string `gorm:"type:varchar(255)"`
	Name           string `gorm:"unique;type:varchar(255)"`
	Cost           uint8
	Detail         string `gorm:"type:text"`
	UnlockCriteria string `gorm:"type:text"`
}
