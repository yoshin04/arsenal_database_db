package queryService

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
)

type SeriesTitleFindManyResult struct {
	SeriesTitles []*string
	TotalLength  int64
}

type ISeriesTitleQueryService interface {
	FindMany() (*SeriesTitleFindManyResult, error)
}

type seriesTitleQueryService struct {
	db *gorm.DB
}

func NewSeriesTitleQueryService(db *gorm.DB) ISeriesTitleQueryService {
	return &seriesTitleQueryService{db}
}

func (s *seriesTitleQueryService) FindMany() (*SeriesTitleFindManyResult, error) {
	log.Println("Running SeriesTitleQueryService.FindMany")
	var modelSeriesTitles []*models.SeriesTitle
	var totalLength int64

	if err := s.db.Model(&models.SeriesTitle{}).Count(&totalLength).Error; err != nil {
		return nil, err
	}

	if err := s.db.Find(&modelSeriesTitles).Error; err != nil {
		return nil, err
	}

	var seriesTitles []*string
	for _, modelSeriesTitle := range modelSeriesTitles {
		seriesTitles = append(seriesTitles, &modelSeriesTitle.Name)
	}

	return &SeriesTitleFindManyResult{
		SeriesTitles: seriesTitles,
		TotalLength:  totalLength,
	}, nil
}
