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

type UrlService struct {
	URLRepository       database.URLRepository
	urlFactory          factory.URLFactory
	urlVisitsRepository database.URLVisitRepository
}

func (urlService *UrlService) GetAll() ([]model.URL, error) {
	return urlService.URLRepository.GetAll()
}

func (urlService *UrlService) Add(newUrl model.NewURL) (model.URL, error) {
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
		retry.WithMaxRetries(3, retry.NewConstant(1*time.Nanosecond)),
		addNewURLWithRetries)

	if err := doRetry; err != nil {
		return model.URL{}, err
	} else {
		return url, err
	}

}

func (urlService *UrlService) GetByExternalID(externalID string) (model.URL, error) {
	return urlService.URLRepository.GetByExternalId(externalID)
}
