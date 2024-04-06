package mocks

import (
	"app/db/models"

	"github.com/stretchr/testify/mock"
)

type MockMSCardRepository struct {
	mock.Mock
}

func (_m *MockMSCardRepository) Upsert(msCard *models.MSCard) error {
	ret := _m.Called(msCard)
	return ret.Error(0)
}
