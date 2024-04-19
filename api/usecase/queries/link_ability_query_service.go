package queryService

import (
	"app/db/models"
	domain "app/domain/card"
	"log"

	"gorm.io/gorm"
)

type LinkAbilityFindManyResult struct {
	LinkAbilities []*domain.LinkAbility
	TotalLength   int64
}

type ILinkAbilityQueryService interface {
	FindMany() (*LinkAbilityFindManyResult, error)
}

type linkAbilityQueryService struct {
	db *gorm.DB
}

func NewLinkAbilityQueryService(db *gorm.DB) ILinkAbilityQueryService {
	return &linkAbilityQueryService{db}
}

func (s *linkAbilityQueryService) FindMany() (*LinkAbilityFindManyResult, error) {
	var modelLinkAbilities []*domain.LinkAbility
	var totalLength int64

	if err := s.db.Model(&models.LinkAbility{}).Count(&totalLength).Error; err != nil {
		log.Printf("Error counting LinkAbilities: %v", err)
		return nil, err
	}

	if err := s.db.Find(&modelLinkAbilities).Error; err != nil {
		log.Printf("Error finding LinkAbilities: %v", err)
		return nil, err
	}

	linkAbilities := make([]*domain.LinkAbility, len(modelLinkAbilities))
	for i, modelLinkAbility := range modelLinkAbilities {
		linkAbilities[i] = &domain.LinkAbility{
			ID:                modelLinkAbility.ID,
			Name:              modelLinkAbility.Name,
			Effect:            modelLinkAbility.Effect,
			RequiredCardCount: modelLinkAbility.RequiredCardCount,
		}
	}

	return &LinkAbilityFindManyResult{
		LinkAbilities: linkAbilities,
		TotalLength:   totalLength,
	}, nil
}
