package mocks

import (
	domain "app/domain/card"

	"github.com/stretchr/testify/mock"
)

type MockLinkAbilityRepository struct {
	mock.Mock
}

func (_m *MockLinkAbilityRepository) FindOrCreate(name string, effect string, requiredCardCount uint8) (*domain.LinkAbility, error) {
	ret := _m.Called(name, effect, requiredCardCount)
	return ret.Get(0).(*domain.LinkAbility), ret.Error(1)
}
