package repository

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IMSCardRepository interface {
	Upsert(card *models.MSCard) error
}

type msCardRepository struct {
	db *gorm.DB
}

func NewMSCardRepository(db *gorm.DB) IMSCardRepository {
	return &msCardRepository{db}
}

func (r *msCardRepository) Upsert(card *models.MSCard) error {
	log.Println("Running MSCardRepository.Upsert")
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"included_code", "no", "image_url", "name", "rarity", "type", "cost", "mobility", "long_range_attack", "close_range_attack", "hp", "total_score", "main_range", "sub_range", "ground_suitability", "space_suitability", "desert_suitability", "underwater_suitability", "sp_attack_name", "sp_cost", "sp_power", "sp_range", "sp_target_type", "sp_detail", "ability_name", "ability_cost", "ability_range", "ability_detail", "series_title", "first_link_ability_id", "second_link_ability_id",
		}),
	}).Create(card).Error
}
