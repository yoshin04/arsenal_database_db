package models

import (
	"gorm.io/gorm"
)

type MSCard struct {
	gorm.Model
	ID                    string `gorm:"primaryKey;unique;type:varchar(255)"`
	IncludedCode          string `gorm:"type:varchar(255)"`
	No                    string `gorm:"type:varchar(255)"`
	ImageURL              string `gorm:"type:varchar(255)"`
	Name                  string `gorm:"type:varchar(255)"`
	Rarity                string `gorm:"type:varchar(8)"`
	Type                  string `gorm:"type:varchar(255)"`
	Cost                  uint8
	Mobility              uint16
	LongRangeAttack       uint16
	CloseRangeAttack      uint16
	HP                    uint16
	TotalScore            uint16
	MainRange             uint16
	SubRange              uint16
	GroundSuitability     string `gorm:"type:varchar(8)"`
	SpaceSuitability      string `gorm:"type:varchar(8)"`
	DesertSuitability     string `gorm:"type:varchar(8)"`
	UnderwaterSuitability string `gorm:"type:varchar(8)"`
	SpAttackName          string `gorm:"type:varchar(255)"`
	SpCost                uint8
	SpPower               uint16
	SpRange               uint8
	SpTargetType          string  `gorm:"type:varchar(255)"`
	SpType                string  `gorm:"type:varchar(255)"`
	SpDetail              string  `gorm:"type:varchar(255)"`
	AbilityName           *string `gorm:"type:varchar(255)"`
	AbilityCost           *uint8
	AbilityRange          *uint8
	AbilityDetail         *string `gorm:"type:varchar(255)"`
	SeriesTitle           string  `gorm:"type:varchar(255)"`

	FirstLinkAbilityID  *string     `gorm:"size:255"`
	SecondLinkAbilityID *string     `gorm:"size:255"`
	FirstLinkAbility    LinkAbility `gorm:"foreignKey:FirstLinkAbilityID;references:ID"`
	SecondLinkAbility   LinkAbility `gorm:"foreignKey:SecondLinkAbilityID;references:ID"`
}
