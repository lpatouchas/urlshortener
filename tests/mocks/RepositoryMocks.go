package mocks

import (
	"github.com/stretchr/testify/mock"
	"urlshortener/model"
)

type MockURLRepository struct {
	mock.Mock
}

func (m *MockURLRepository) Add(url model.URL) (model.URL, error) {
	args := m.Called(url)
	return args.Get(0).(model.URL), args.Error(1)
}

func (m *MockURLRepository) GetAll() ([]model.URL, error) {
	args := m.Called()
	return args.Get(0).([]model.URL), args.Error(1)
}

func (m *MockURLRepository) GetByExternalId(externalID string) (model.URL, error) {
	args := m.Called(externalID)
	return args.Get(0).(model.URL), args.Error(1)
}

type MockURLVisitRepository struct {
	mock.Mock
}

func (m *MockURLVisitRepository) Add(urlVisit model.URLVisit) error {
	args := m.Called(urlVisit)
	return args.Error(0)
}

func (m *MockURLVisitRepository) CountURLVisits(url model.URL) (int, error) {
	args := m.Called(url)
	return args.Int(0), args.Error(1)
}
