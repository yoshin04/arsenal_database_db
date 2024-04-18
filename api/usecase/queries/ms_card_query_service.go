package queryService

import (
	"app/db/models"
	domain "app/domain/card"
	"log"

	"gorm.io/gorm"
)

type MsCardFindManyInput struct {
	Offset                  int
	Limit                   int
	Costs                   []uint8
	Rarities                []string
	MsTypes                 []string
	GroundSuitabilities     []string
	SpaceSuitabilities      []string
	DesertSuitabilities     []string
	UnderwaterSuitabilities []string
	IncludedCode            *string
	LinkAbilityID           *string
	Keyword                 *string
}
type MsCardFindManyResult struct {
	Cards       []*domain.MsCard
	TotalLength int64
}

type IMsCardQueryService interface {
	FindMany(input MsCardFindManyInput) (*MsCardFindManyResult, error)
	FindOneById(id string) (*domain.MsCard, error)
}

type msCardQueryService struct {
	db *gorm.DB
}

func NewMsCardQueryService(db *gorm.DB) IMsCardQueryService {
	return &msCardQueryService{db}
}

func (s *msCardQueryService) FindMany(input MsCardFindManyInput) (*MsCardFindManyResult, error) {
	log.Println("Running MsCardQueryService.FindMany")
	var modelMsCards []*models.MSCard
	var totalLength int64

	baseQuery := s.db.Model(&models.MSCard{})

	if len(input.Costs) > 0 {
		costsInt := make([]interface{}, len(input.Costs))
		for i, cost := range input.Costs {
			costsInt[i] = int(cost)
		}
		if input.Costs[len(input.Costs)-1] == 7 {
			baseQuery = baseQuery.Where("cost >= ?", 7)
		} else {
			baseQuery = baseQuery.Where("cost IN (?)", costsInt) // パラメータの部分を修正
		}
	}

	if len(input.Rarities) > 0 {
		baseQuery = baseQuery.Where("rarity IN ?", input.Rarities)
	}
	if len(input.MsTypes) > 0 {
		baseQuery = baseQuery.Where("type In ?", input.MsTypes)
	}
	if len(input.GroundSuitabilities) > 0 {
		baseQuery = baseQuery.Where("ground_suitability IN ?", input.GroundSuitabilities)
	}
	if len(input.SpaceSuitabilities) > 0 {
		baseQuery = baseQuery.Where("space_suitability IN ?", input.SpaceSuitabilities)
	}
	if len(input.DesertSuitabilities) > 0 {
		baseQuery = baseQuery.Where("desert_suitability IN ?", input.DesertSuitabilities)
	}
	if len(input.UnderwaterSuitabilities) > 0 {
		baseQuery = baseQuery.Where("underwater_suitability IN ?", input.UnderwaterSuitabilities)
	}
	if input.IncludedCode != nil && *input.IncludedCode != "" {
		baseQuery = baseQuery.Where("included_code = ?", input.IncludedCode)
	}
	if input.LinkAbilityID != nil && *input.LinkAbilityID != "" {
		baseQuery = baseQuery.Where("first_link_ability_id = ? OR second_link_ability_id = ?", input.LinkAbilityID, input.LinkAbilityID)
	}

	if input.Keyword != nil && *input.Keyword != "" {
		keyword := "%" + *input.Keyword + "%"
		// 収録コード、カードNo、カード名を連結して部分一致検索（スペースで区切る）
		baseQuery = baseQuery.Where("CONCAT(included_code, ' ', no, ' ', name) LIKE ?", keyword)
	}

	if err := baseQuery.Count(&totalLength).Error; err != nil {
		return &MsCardFindManyResult{}, err
	}

	result := baseQuery.Offset(input.Offset).Limit(input.Limit).Preload("FirstLinkAbility").Preload("SecondLinkAbility").Find(&modelMsCards)
	if result.Error != nil {
		return &MsCardFindManyResult{}, result.Error
	}

	var domainMsCards []*domain.MsCard
	for _, modelMsCard := range modelMsCards {
		if modelMsCard != nil {
			log.Printf("modelMsCard: %v", modelMsCard)
			domainMsCard := domain.ToDomainMsCard(modelMsCard)
			domainMsCards = append(domainMsCards, domainMsCard)
		}
	}
	return &MsCardFindManyResult{
		Cards:       domainMsCards,
		TotalLength: totalLength,
	}, nil
}

func (s *msCardQueryService) FindOneById(id string) (*domain.MsCard, error) {
	log.Printf("Running MsCardQueryService.FindOneById: %v", id)
	var modelMsCard models.MSCard
	result := s.db.Preload("FirstLinkAbility").Preload("SecondLinkAbility").First(&modelMsCard, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return domain.ToDomainMsCard(&modelMsCard), nil
}
