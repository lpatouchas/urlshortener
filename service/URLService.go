package service

import (
	"context"
	"errors"
	"github.com/sethvargo/go-retry"
	"log"
	"time"
	"urlshortener/database"
	"urlshortener/factory"
	"urlshortener/model"
)

type UrlServiceImpl struct {
	URLRepository       database.DBRepository[model.URL]
	urlFactory          factory.URLFactory
	urlVisitsRepository database.URLVisitRepository
}

func NewUrlService(urlRepository database.DBRepository[model.URL], urlVisitRepository database.URLVisitRepository) UrlServiceImpl {
	return UrlServiceImpl{
		URLRepository:       urlRepository,
		urlVisitsRepository: urlVisitRepository,
	}
}

func (urlService *UrlServiceImpl) GetAll() ([]model.URL, error) {
	return urlService.URLRepository.GetAll()
}

func (urlService *UrlServiceImpl) Add(newUrl model.NewURL) (model.URL, error) {
	urlToAdd := urlService.urlFactory.FromNewURL(newUrl)

	//Prepare retriable function in case of Unique constraint violation on generatedExternalId
	retryCount := 0
	var url model.URL
	addNewURLWithRetries := func(ctx context.Context) error {
		added, err := urlService.URLRepository.Add(urlToAdd)
		url = added //TODO code smell
		retryCount++
		if err != nil {
			if errors.Is(database.ErrUniqueConstraint, err) {
				log.Printf("error during add, retrying (attempt %d)", retryCount)
				return retry.RetryableError(err)
			} else {
				return err
			}
		}
		return nil
	}

	doRetry := retry.Do(context.Background(),
		retry.WithMaxRetries(2, retry.NewConstant(1*time.Nanosecond)),
		addNewURLWithRetries)

	if err := doRetry; err != nil {
		return model.URL{}, err
	} else {
		return url, err
	}

}

func (urlService *UrlServiceImpl) GetByExternalID(externalID string) (model.URL, error) {
	return urlService.URLRepository.GetByExternalId(externalID)
}
