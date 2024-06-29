package database

import (
	"database/sql"
	"errors"
	"log"
	"time"
	"urlshortener/factory"
	"urlshortener/model"
	"urlshortener/urlErrors"
)

type URLVisitRepositoryImpl struct {
	urlVisitFactory factory.URLVisitFactory
}

type URLVisitRepository interface {
	Add(urlVisit model.URLVisit) error
	CountURLVisits(url model.URL) (int, error)
}

const (
	insertOneVisit    = "INSERT INTO public.short_url_visits (short_url_id, visited_at) VALUES($1,$2) RETURNING id"
	countVisitsForOne = "select count(suv.id) from short_url_visits suv inner join short_urls su on suv.short_url_id = su.id \nwhere su.externalid = $1"
)

func (repo *URLVisitRepositoryImpl) Add(urlVisit model.URLVisit) error {
	var newID int
	err := Db.QueryRow(insertOneVisit, urlVisit.URL.ID, time.Now()).Scan(&newID)
	if err != nil {
		log.Printf("An error occured during persisting url visit %v with error: %v\n", urlVisit, err)
	}
	log.Printf("New URLVisit added succesfully with id %d", newID)
	return err
}

func (repo *URLVisitRepositoryImpl) CountURLVisits(url model.URL) (int, error) {
	var count int
	err := Db.QueryRow(countVisitsForOne, url.ExternalID).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, urlErrors.FromExternalIDWithCustomMessageAndCode(url.ExternalID, "Could not find visits for this URL external ID", 404)

	}
	if err != nil {
		log.Printf("Error querying visit count for externalID %s: %v", url.ExternalID, err)
		return 0, err
	}
	return count, nil
}
