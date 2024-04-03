package models

import "gorm.io/gorm"

type PLCard struct {
	gorm.Model
	ID               string `gorm:"primaryKey;unique;type:varchar(255)"`
	IncludeCode      string `gorm:"type:varchar(255)"`
	No               string `gorm:"type:varchar(255)"`
	ImageURL         string `gorm:"type:varchar(255)"`
	Name             string `gorm:"type:varchar(255)"`
	Rarity           string `gorm:"type:varchar(8)"`
	Type             string `gorm:"type:varchar(255)"`
	Cost             uint8
	Mobility         uint16
	LongRangeAttack  uint16
	CloseRangeAttack uint16
	HP               uint16
	TotalScore       uint16
	PlSkillName      string `gorm:"type:varchar(255)"`
	PlSkillCondition string `gorm:"type:varchar(255)"`
	PlSkillDetail    string `gorm:"type:varchar(255)"`
	SeriesTitle      string `gorm:"type:varchar(255)"`

	FirstLinkAbilityID  *string     `gorm:"size:255"`
	SecondLinkAbilityID *string     `gorm:"size:255"`
	FirstLinkAbility    LinkAbility `gorm:"foreignKey:FirstLinkAbilityID;references:ID"`
	SecondLinkAbility   LinkAbility `gorm:"foreignKey:SecondLinkAbilityID;references:ID"`
}
