package repository

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ISeriesTitleRepository interface {
	Upsert(name string) error
}

type seriesTitleRepository struct {
	db *gorm.DB
}

func NewSeriesTitleRepository(db *gorm.DB) ISeriesTitleRepository {
	return &seriesTitleRepository{db}
}

func (r *seriesTitleRepository) Upsert(name string) error {
	log.Println("Running SeriesTitleRepository.Upsert")
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{}),
	}).Create(&models.SeriesTitle{Name: name}).Error
}
