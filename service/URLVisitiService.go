package service

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
	"urlshortener/database"
	"urlshortener/factory"
	"urlshortener/model"
)

type URLVisitService struct {
	urlService         GetByExternalID
	urlVisitRepository database.URLVisitRepository
	urlVisitFactory    factory.URLVisitFactory
	redirectCache      *cache.Cache
}

func NewURLVisitService(urlService GetByExternalID, urlVisitRepository database.URLVisitRepository, urlVisitFactory factory.URLVisitFactory) URLVisitService {
	return URLVisitService{
		urlService:         urlService,
		urlVisitRepository: urlVisitRepository,
		urlVisitFactory:    urlVisitFactory,
		redirectCache:      cache.New(6*time.Hour, 12*time.Hour), // Adjust the expiration times as needed
	}
}

func (urlVisitService *URLVisitService) GetRedirectURL(urlExternalId string) (string, error) {

	// Check cache first
	if cachedURL, found := urlVisitService.redirectCache.Get(urlExternalId); found {
		log.Printf("Cache hit for external ID: %s", urlExternalId)
		urlVisitService.registerVisit(cachedURL.(model.URL))
		return cachedURL.(model.URL).LongURL, nil
	}

	url, err := urlVisitService.urlService.GetByExternalID(urlExternalId)

	if err != nil {
		log.Printf("Could not load ULR with external id %s, due to error: %v\n", urlExternalId, err)
		return "", err
	}

	urlVisitService.registerVisit(url)

	urlVisitService.redirectCache.Set(urlExternalId, url, cache.DefaultExpiration)
	return url.LongURL, nil
}

func (urlVisitService *URLVisitService) visitURL(url model.URL) error {
	log.Printf("Registered new URL visit for URL %v", url)
	return urlVisitService.urlVisitRepository.Add(urlVisitService.urlVisitFactory.FromURL(url))
}

func (urlVisitService *URLVisitService) CountURLVisits(externalID string) (model.URLVisits, error) {
	url, err := urlVisitService.urlService.GetByExternalID(externalID)
	if err != nil {
		//TOOD custom error and logging
		return model.URLVisits{}, err
	}
	count, err := urlVisitService.urlVisitRepository.CountURLVisits(url)

	return model.URLVisits{ExternalID: externalID, Visits: count}, err
}

func (urlVisitService *URLVisitService) registerVisit(url model.URL) {
	go func() {
		err := urlVisitService.visitURL(url)
		if err != nil {
			log.Printf("Could not register the visit of URL %v due to error: %v", url, err)
		}
	}()
}
