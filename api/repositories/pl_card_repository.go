package repository

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPLCardRepository interface {
	Upsert(card *models.PLCard) error
}

type plCardRepository struct {
	db *gorm.DB
}

func NewPLCardRepository(db *gorm.DB) IPLCardRepository {
	return &plCardRepository{db}
}

func (r *plCardRepository) Upsert(card *models.PLCard) error {
	log.Println("Running PLCardRepository.Upsert")
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"include_code", "no", "image_url", "name", "rarity", "type", "cost", "mobility", "long_range_attack", "close_range_attack", "hp", "total_score", "pl_skill_name", "pl_skill_condition", "pl_skill_detail", "series_title", "first_link_ability_id", "second_link_ability_id"}),
	}).Create(card).Error
}
