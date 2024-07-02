package mocks

import (
	"github.com/stretchr/testify/mock"
	"urlshortener/model"
)

type MockURLVisitService struct {
	mock.Mock
}

func (m *MockURLVisitService) GetRedirectURL(urlExternalId string) (string, error) {
	args := m.Called(urlExternalId)
	return args.Get(0).(string), args.Error(1)
}

type MockGetByExternalID struct {
	mock.Mock
}

func (m *MockGetByExternalID) GetByExternalID(externalID string) (model.URL, error) {
	args := m.Called(externalID)
	return args.Get(0).(model.URL), args.Error(1)
}
