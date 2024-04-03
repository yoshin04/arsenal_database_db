package models

import "gorm.io/gorm"

type GameDeck struct {
	gorm.Model
	ID                   string        `gorm:"primaryKey;unique;type:varchar(255)"`
	Name                 string        `gorm:"type:varchar(255)"`
	UserId               string        `gorm:"type:varchar(255)"`
	User                 User          `gorm:"foreignKey:UserId;references:ID"`
	FirstMSCardID        *string       `gorm:"type:varchar(255)"`
	FirstMSCard          *MSCard       `gorm:"foreignKey:FirstMSCardID;references:ID"`
	SecondMSCardID       *string       `gorm:"type:varchar(255)"`
	SecondMSCard         *MSCard       `gorm:"foreignKey:SecondMSCardID;references:ID"`
	ThirdMSCardID        *string       `gorm:"type:varchar(255)"`
	ThirdMSCard          *MSCard       `gorm:"foreignKey:ThirdMSCardID;references:ID"`
	FourthMSCardID       *string       `gorm:"type:varchar(255)"`
	FourthMSCard         *MSCard       `gorm:"foreignKey:FourthMSCardID;references:ID"`
	FifthMSCardID        *string       `gorm:"type:varchar(255)"`
	FifthMSCard          *MSCard       `gorm:"foreignKey:FifthMSCardID;references:ID"`
	FistPLCardID         *string       `gorm:"type:varchar(255)"`
	FistPLCard           *PLCard       `gorm:"foreignKey:FistPLCardID;references:ID"`
	SecondPLCardID       *string       `gorm:"type:varchar(255)"`
	SecondPLCard         *PLCard       `gorm:"foreignKey:SecondPLCardID;references:ID"`
	ThirdPLCardID        *string       `gorm:"type:varchar(255)"`
	ThirdPLCard          *PLCard       `gorm:"foreignKey:ThirdPLCardID;references:ID"`
	FourthPLCardID       *string       `gorm:"type:varchar(255)"`
	FourthPLCard         *PLCard       `gorm:"foreignKey:FourthPLCardID;references:ID"`
	FifthPLCardID        *string       `gorm:"type:varchar(255)"`
	FifthPLCard          *PLCard       `gorm:"foreignKey:FifthPLCardID;references:ID"`
	FistTacticalCardID   *string       `gorm:"type:varchar(255)"`
	FistTacticalCard     *TacticalCard `gorm:"foreignKey:FistTacticalCardID;references:ID"`
	SecondTacticalCardID *string       `gorm:"type:varchar(255)"`
	SecondTacticalCard   *TacticalCard `gorm:"foreignKey:SecondTacticalCardID;references:ID"`
	ThirdTacticalCardID  *string       `gorm:"type:varchar(255)"`
	ThirdTacticalCard    *TacticalCard `gorm:"foreignKey:ThirdTacticalCardID;references:ID"`
}
