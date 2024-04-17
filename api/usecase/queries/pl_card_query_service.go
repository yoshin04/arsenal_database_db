package queryService

import (
	"app/db/models"
	domain "app/domain/card"
	"log"

	"gorm.io/gorm"
)

type PlCardFindManyInput struct {
	Offset        int
	Limit         int
	Costs         []uint8
	Rarities      []string
	PlTypes       []string
	IncludedCode  *string
	LinkAbilityID *string
	Keyword       *string
}

type PlCardFindManyResult struct {
	Cards       []*domain.PlCard
	TotalLength int64
}

type IPlCardQueryService interface {
	FindMany(input PlCardFindManyInput) (*PlCardFindManyResult, error)
	// FindOneById(id string) (*domain.PlCard, error)
}

type plCardQueryService struct {
	db *gorm.DB
}

func NewPlCardQueryService(db *gorm.DB) IPlCardQueryService {
	return &plCardQueryService{db}
}

func (s *plCardQueryService) FindMany(input PlCardFindManyInput) (*PlCardFindManyResult, error) {
	log.Println("Running PlCardQueryService.FindMany")
	var modelPlCards []*models.PLCard
	var totalLength int64

	baseQuery := s.db.Model(&models.PLCard{})

	if len(input.Costs) > 0 {
		costsInt := make([]interface{}, len(input.Costs))
		for i, cost := range input.Costs {
			costsInt[i] = int(cost)
		}
		if input.Costs[len(input.Costs)-1] == 7 {
			baseQuery = baseQuery.Where("cost >= ?", 7)
		} else {
			baseQuery = baseQuery.Where("cost IN (?)", costsInt)
		}
	}

	if len(input.Rarities) > 0 {
		baseQuery = baseQuery.Where("rarity IN (?)", input.Rarities)
	}
	if len(input.PlTypes) > 0 {
		baseQuery = baseQuery.Where("type IN (?)", input.PlTypes)
	}
	if input.IncludedCode != nil {
		baseQuery = baseQuery.Where("include_code = ?", *input.IncludedCode)
	}
	if input.LinkAbilityID != nil && *input.LinkAbilityID != "" {
		baseQuery = baseQuery.Where("first_link_ability_id = ? OR second_link_ability_id = ?", *input.LinkAbilityID, *input.LinkAbilityID)
	}

	if input.Keyword != nil && *input.Keyword != "" {
		keyword := "%" + *input.Keyword + "%"
		baseQuery = baseQuery.Where("CONCAT(include_code, ' ', no, ' ', name) LIKE ?", keyword)
	}

	if err := baseQuery.Count(&totalLength).Error; err != nil {
		return &PlCardFindManyResult{}, err
	}

	result := baseQuery.Offset(input.Offset).Limit(input.Limit).Preload("FirstLinkAbility").Preload("SecondLinkAbility").Find(&modelPlCards)
	if result.Error != nil {
		return &PlCardFindManyResult{}, result.Error

	}

	var domainPlCards []*domain.PlCard
	for _, modelPlCard := range modelPlCards {
		if modelPlCard != nil {
			log.Printf("modelPlCard: %v\n", modelPlCard)
			domainPlCard := domain.ToDomainPlCard(modelPlCard)
			domainPlCards = append(domainPlCards, domainPlCard)
		}
	}

	return &PlCardFindManyResult{
		Cards:       domainPlCards,
		TotalLength: totalLength,
	}, nil
}

// func (s *plCardQueryService) FindOneById(id string) (*domain.PlCard, error) {
// 	var plCards = mocks.GenerateDummyPlCards()
// 	for _, plCard := range plCards {
// 		if plCard.Id == id {
// 			return plCard, nil
// 		}
// 	}
// 	return nil, nil
// }
