package service

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"urlshortener/database"
	"urlshortener/factory"
	"urlshortener/model"
	"urlshortener/tests/mocks"
	"urlshortener/urlErrors"
)

// Trying table driven tests
func TestUrlServiceImpl_Add(t *testing.T) {
	err := godotenv.Load("../tests/.env")
	if err != nil {
		println("error loading test envs")
	}
	type fields struct {
		URLRepository       *mocks.MockURLRepository
		urlFactory          factory.URLFactory
		urlVisitsRepository *mocks.MockURLVisitRepository
	}
	type args struct {
		newUrl model.NewURL
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		mock    func(f fields)
		want    model.URL
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Add OK",
			fields: fields{
				URLRepository:       new(mocks.MockURLRepository),
				urlFactory:          factory.URLFactory{},
				urlVisitsRepository: new(mocks.MockURLVisitRepository),
			},
			args: args{
				newUrl: model.NewURL{LongURL: "http://www.google.com"},
			},
			mock: func(f fields) {
				expectedURL := model.URL{ID: 1, ExternalID: "abc123", LongURL: "http://www.google.com"}
				f.URLRepository.On("Add", mock.MatchedBy(func(url model.URL) bool {
					return url.LongURL == "http://www.google.com"
				})).Return(expectedURL, nil).Once()
			},
			want: model.URL{ID: 1, ExternalID: "abc123", LongURL: "http://www.google.com"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "Add Unique Constraint Error",
			fields: fields{
				URLRepository:       new(mocks.MockURLRepository),
				urlFactory:          factory.URLFactory{},
				urlVisitsRepository: new(mocks.MockURLVisitRepository),
			},
			args: args{
				newUrl: model.NewURL{LongURL: "http://www.google.com"},
			},
			mock: func(f fields) {
				f.URLRepository.On("Add", mock.MatchedBy(func(url model.URL) bool {
					return url.LongURL == "http://www.google.com"
				})).Return(model.URL{}, database.ErrUniqueConstraint).Times(3)
			},
			want: model.URL{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.True(t, errors.Is(err, database.ErrUniqueConstraint), "expected ErrUniqueConstraint error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize mocks and the urlService
			tt.mock(tt.fields)
			urlService := &UrlServiceImpl{
				URLRepository:       tt.fields.URLRepository,
				urlFactory:          tt.fields.urlFactory,
				urlVisitsRepository: tt.fields.urlVisitsRepository,
			}
			got, err := urlService.Add(tt.args.newUrl)
			if !tt.wantErr(t, err, fmt.Sprintf("Add(%v)", tt.args.newUrl)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Add(%v)", tt.args.newUrl)
			tt.fields.URLRepository.AssertExpectations(t)
		})
	}
}

// useless test but just checking stuff
func TestUrlServiceImpl_GetByExternalID(t *testing.T) {
	type fields struct {
		URLRepository       *mocks.MockURLRepository
		urlFactory          factory.URLFactory
		urlVisitsRepository database.URLVisitRepository
	}
	type args struct {
		externalID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(f fields)
		want    model.URL
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test ExternalIdNotFound",
			fields: fields{
				URLRepository:       new(mocks.MockURLRepository),
				urlFactory:          factory.URLFactory{},
				urlVisitsRepository: new(mocks.MockURLVisitRepository),
			},
			args: args{externalID: "xxxxx"},
			mock: func(f fields) {
				f.URLRepository.On("GetByExternalId", "xxxxx").Return(model.URL{}, urlErrors.FromExternalID("xxxxx")).Once()
			},
			want: model.URL{},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				if assert.Error(t, err) {
					var redirectErr *urlErrors.RedirectError
					if errors.As(err, &redirectErr) {
						assert.Equal(t, 400, redirectErr.StatusCode)
						assert.Equal(t, "something went wrong", redirectErr.Err.Error())
						assert.Equal(t, "xxxxx", redirectErr.UrlExternalId)
						return true
					}
				}
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields)
			urlService := &UrlServiceImpl{
				URLRepository:       tt.fields.URLRepository,
				urlFactory:          tt.fields.urlFactory,
				urlVisitsRepository: tt.fields.urlVisitsRepository,
			}
			got, err := urlService.GetByExternalID(tt.args.externalID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByExternalID(%v)", tt.args.externalID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetByExternalID(%v)", tt.args.externalID)
		})
	}
}
