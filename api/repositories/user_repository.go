package repository

import (
	"app/db/models"
	"log"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(card *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	log.Println("Running UserRepository.Create")
	if err := r.db.Create(user).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return err
	}
	return nil
}
