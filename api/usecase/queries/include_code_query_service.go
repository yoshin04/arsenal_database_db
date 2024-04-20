package queryService

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
)

type IncludeCodeFindManyResult struct {
	IncludeCodes []*string
	TotalLength  int64
}

type IIncludeCodeQueryService interface {
	FindMany() (*IncludeCodeFindManyResult, error)
}

type includeCodeQueryService struct {
	db *gorm.DB
}

func NewIncludeCodeQueryService(db *gorm.DB) IIncludeCodeQueryService {
	return &includeCodeQueryService{db}
}

func (s *includeCodeQueryService) FindMany() (*IncludeCodeFindManyResult, error) {
	log.Println("Running IncludeCodeQueryService.FindMany")
	var modelIncludeCodes []*models.IncludeCode
	var totalLength int64

	if err := s.db.Model(&models.IncludeCode{}).Count(&totalLength).Error; err != nil {
		log.Printf("Error counting IncludeCodes: %v", err)
		return nil, err
	}

	if err := s.db.Find(&modelIncludeCodes).Error; err != nil {
		return nil, err
	}

	var includeCodes []*string
	for _, modelIncludeCode := range modelIncludeCodes {
		includeCodes = append(includeCodes, &modelIncludeCode.Name)
	}

	return &IncludeCodeFindManyResult{
		IncludeCodes: includeCodes,
		TotalLength:  totalLength,
	}, nil
}
