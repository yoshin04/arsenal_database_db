package repository

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITacticalCardRepository interface {
	Upsert(card *models.TacticalCard) error
}

type tacticalRepository struct {
	db *gorm.DB
}

func NewTacticalCardRepository(db *gorm.DB) ITacticalCardRepository {
	return &tacticalRepository{db}
}

func (r *tacticalRepository) Upsert(card *models.TacticalCard) error {
	log.Println("Running TacticalCardRepository.Upsert")

	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"image_url", "name", "cost", "detail", "unlock_criteria",
		}),
	}).Create(card).Error
}
