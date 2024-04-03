package models

import "gorm.io/gorm"

type LinkAbility struct {
	gorm.Model
	ID                string `gorm:"primaryKey;unique;type:varchar(255)"`
	Name              string `gorm:"type:varchar(255);unique"`
	RequiredCardCount uint8  `gorm:"type:tinyint"`
	Effect            string `gorm:"type:text"`
}
