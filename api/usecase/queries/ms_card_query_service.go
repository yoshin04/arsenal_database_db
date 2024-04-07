package queryService

import (
	"app/db/models"
	domain "app/domain/card"
	"log"

	"gorm.io/gorm"
)

type MsCardFindManyInput struct {
	Offset int
	Limit  int
}

type IMsCardQueryService interface {
	FindMany(input MsCardFindManyInput) ([]*domain.MsCard, error)
}

type msCardQueryService struct {
	db *gorm.DB
}

func NewMsCardQueryService(db *gorm.DB) IMsCardQueryService {
	return &msCardQueryService{db}
}

func (s *msCardQueryService) FindMany(input MsCardFindManyInput) ([]*domain.MsCard, error) {
	log.Println("Running MsCardQueryService.FindMany")
	var modelMsCards []*models.MSCard
	result := s.db.Offset(input.Offset).Limit(input.Limit).Preload("FirstLinkAbility").Preload("SecondLinkAbility").Find(&modelMsCards)
	if result.Error != nil {
		return nil, result.Error
	}

	var domainMsCards []*domain.MsCard
	for _, modelMsCard := range modelMsCards {
		if modelMsCard != nil {
			log.Printf("modelMsCard: %v", modelMsCard)
			domainMsCard := domain.ToDomainMsCard(modelMsCard)
			domainMsCards = append(domainMsCards, domainMsCard)
		}
	}
	return domainMsCards, result.Error
}
