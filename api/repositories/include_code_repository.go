package repository

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IIncludeCodeRepository interface {
	Upsert(name string) error
}

type includeCodeRepository struct {
	db *gorm.DB
}

func NewIncludeCodeRepository(db *gorm.DB) IIncludeCodeRepository {
	return &includeCodeRepository{db}
}

func (r *includeCodeRepository) Upsert(name string) error {
	log.Println("Running IncludeCodeRepository.Upsert")
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{}),
	}).Create(&models.IncludeCode{Name: name}).Error
}
