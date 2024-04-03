package repository

import (
	"app/db/models"
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"gorm.io/gorm"
)

type ILinkAbilityRepository interface {
	FindOrCreate(name string, effect string, requiredCardCount uint8) (models.LinkAbility, error)
}

type LinkAbilityRepository struct {
	db *gorm.DB
}

func NewLinkAbilityRepository(db *gorm.DB) ILinkAbilityRepository {
	return &LinkAbilityRepository{db}
}

func (r *LinkAbilityRepository) FindOrCreate(name string, effect string, requiredCardCount uint8) (models.LinkAbility, error) {
	log.Println("Running LinkAbilityRepository.FindOrCreate")
	var linkAbility models.LinkAbility

	if name == "-" {
		return linkAbility, nil
	}

	err := r.db.Where(&models.LinkAbility{Name: name}).First(&linkAbility).Error
	if err == nil {
		return linkAbility, nil
	}

	if err == gorm.ErrRecordNotFound {
		id, err := gonanoid.New()
		if err != nil {
			return models.LinkAbility{}, err
		}

		linkAbility = models.LinkAbility{
			ID:                id,
			Name:              name,
			Effect:            effect,
			RequiredCardCount: requiredCardCount,
		}
		err = r.db.Create(&linkAbility).Error
		if err != nil {
			return models.LinkAbility{}, err
		}
	}

	return linkAbility, nil

}
