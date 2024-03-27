package queryService

import (
	domain "app/domain/card"
	"app/mock"
)

type IPlCardQueryService interface {
	FindAll() ([]*domain.PlCard, error)
	FindOneById(id string) (*domain.PlCard, error)
}

type plCardQueryService struct{}

func NewPlCardQueryService() IPlCardQueryService {
	return &plCardQueryService{}
}

func (s *plCardQueryService) FindAll() ([]*domain.PlCard, error) {
	var plCards = mock.GenerateDummyPlCards()
	return plCards, nil
}

func (s *plCardQueryService) FindOneById(id string) (*domain.PlCard, error) {
	var plCards = mock.GenerateDummyPlCards()
	for _, plCard := range plCards {
		if plCard.Id == id {
			return plCard, nil
		}
	}
	return nil, nil
}
