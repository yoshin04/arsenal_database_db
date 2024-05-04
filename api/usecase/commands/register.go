package usecase

import (
	"app/db/models"
	domain "app/domain/user"
	repository "app/repositories"
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type RegisterUsecaseInput struct {
	Name        string
	FirebaseUid string
}

type IRegisterUsecase interface {
	Run(input RegisterUsecaseInput) (domain.User, error)
}

type registerUsecase struct {
	userRepo repository.IUserRepository
}

func NewRegisterUsecase(userRepo repository.IUserRepository) IRegisterUsecase {
	return &registerUsecase{
		userRepo: userRepo,
	}
}

func (uc *registerUsecase) Run(input RegisterUsecaseInput) (domain.User, error) {
	log.Printf("Running register usecase with params: Name=%s, FirebaseUID=%s", input.Name, input.FirebaseUid)
	id, err := gonanoid.New()
	if err != nil {
		return domain.User{}, err
	}

	userModel := models.User{
		ID:          id,
		Name:        input.Name,
		FirebaseUID: input.FirebaseUid,
	}

	if err := uc.userRepo.Create(&userModel); err != nil {
		log.Printf("Error Create user: %v", err)
		return domain.User{}, err
	}

	return domain.User{
		Id:        userModel.ID,
		Name:      userModel.Name,
		AvatarUrl: nil,
	}, nil
}
