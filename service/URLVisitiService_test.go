package service

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"urlshortener/factory"
	"urlshortener/model"
	mocks2 "urlshortener/tests/mocks"
)

// test cache is being utilized
func TestURLVisitService_GetRedirectURL(t *testing.T) {
	type fields struct {
		urlService         *mocks2.MockGetByExternalID
		urlVisitRepository *mocks2.MockURLVisitRepository
		urlVisitFactory    factory.URLVisitFactory
		redirectCache      *cache.Cache
	}
	type args struct {
		urlExternalId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(f fields)
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Multiple calls, only one reaches urlService",
			fields: fields{
				urlService:         new(mocks2.MockGetByExternalID),
				urlVisitRepository: new(mocks2.MockURLVisitRepository),
				urlVisitFactory:    factory.URLVisitFactory{},
				redirectCache:      cache.New(1*time.Minute, 1*time.Minute),
			},
			args: args{
				urlExternalId: "externalID",
			},
			mock: func(f fields) {
				expectedURL := model.URL{ExternalID: "externalID", LongURL: "http://example.com"}
				f.urlService.On("GetByExternalID", "externalID").Return(expectedURL, nil).Once()
			},
			want:    "http://example.com",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlVisitService := &URLVisitService{
				urlService:         tt.fields.urlService,
				urlVisitRepository: tt.fields.urlVisitRepository,
				urlVisitFactory:    tt.fields.urlVisitFactory,
				redirectCache:      tt.fields.redirectCache,
			}

			tt.mock(tt.fields)

			var got string
			var err error
			for i := 0; i < 3; i++ { // Make multiple calls
				got, err = urlVisitService.GetRedirectURL(tt.args.urlExternalId)
			}

			if !tt.wantErr(t, err, fmt.Sprintf("GetRedirectURL(%v)", tt.args.urlExternalId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetRedirectURL(%v)", tt.args.urlExternalId)

			// Assert that GetByExternalID is called only once
			tt.fields.urlService.AssertNumberOfCalls(t, "GetByExternalID", 1)
			tt.fields.urlService.AssertExpectations(t)
		})
	}
}
