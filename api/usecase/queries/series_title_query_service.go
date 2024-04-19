package queryService

import (
	"app/db/models"

	"gorm.io/gorm"
)

type SeriesTitleFindManyResult struct {
	SeriesTitles []*string
	TotalLength  int64
}

type ISeriesTitleQueryService interface {
	FindAll() (*SeriesTitleFindManyResult, error)
}

type seriesTitleQueryService struct {
	db *gorm.DB
}

func NewSeriesTitleQueryService(db *gorm.DB) ISeriesTitleQueryService {
	return &seriesTitleQueryService{db}
}

func (s *seriesTitleQueryService) FindAll() (*SeriesTitleFindManyResult, error) {
	var modelSeriesTitles []*models.SeriesTitle
	var totalLength int64

	// 先に総数を取得
	if err := s.db.Model(&models.SeriesTitle{}).Count(&totalLength).Error; err != nil {
		return nil, err
	}

	// シリーズタイトルのリストを取得
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
