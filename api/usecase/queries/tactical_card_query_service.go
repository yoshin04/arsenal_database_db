package queryService

import (
	"app/db/models"
	domain "app/domain/tactical_card"
	"log"

	"gorm.io/gorm"
)

type TacticalCardFindManyInput struct {
	Costs []uint8
}

type TacticalCardFindManyResult struct {
	Cards       []*domain.TacticalCard
	TotalLength int64
}

type ITacticalCardQueryService interface {
	FindAll(input TacticalCardFindManyInput) (*TacticalCardFindManyResult, error)
}

type tacticalCardQueryService struct {
	db *gorm.DB
}

func NewTacticalCardQueryService(db *gorm.DB) ITacticalCardQueryService {
	return &tacticalCardQueryService{db}
}

func (s *tacticalCardQueryService) FindAll(input TacticalCardFindManyInput) (*TacticalCardFindManyResult, error) {
	log.Println("Running TacticalCardQueryService.FindAll")
	var modelTacticalCards []*models.TacticalCard
	var totalLength int64

	baseQuery := s.db.Model(&models.TacticalCard{})

	if len(input.Costs) > 0 {
		costsInt := make([]interface{}, len(input.Costs))
		for i, cost := range input.Costs {
			costsInt[i] = int(cost)
		}
		baseQuery = baseQuery.Where("cost IN (?)", costsInt)
	}

	if err := baseQuery.Count(&totalLength).Error; err != nil {
		return &TacticalCardFindManyResult{}, err
	}

	result := baseQuery.Find(&modelTacticalCards)
	if result.Error != nil {
		return &TacticalCardFindManyResult{}, result.Error
	}

	var domainTacticalCards []*domain.TacticalCard
	for _, modelTacticalCard := range modelTacticalCards {
		if modelTacticalCard != nil {
			domainTacticalCard := domain.ToDomainTacticalCard(modelTacticalCard)
			domainTacticalCards = append(domainTacticalCards, domainTacticalCard)
		}
	}

	return &TacticalCardFindManyResult{
		Cards:       domainTacticalCards,
		TotalLength: totalLength,
	}, nil
}
