package models

import "gorm.io/gorm"

type IncludeCode struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);uniqueIndex"`
}
