package database

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"log"
	"time"
	"urlshortener/factory"
	"urlshortener/model"
	"urlshortener/urlErrors"
)

type URLRepository struct {
	urlFactory factory.URLFactory
}

const (
	UniqueConstraintViolation = "23505"

	GetAll          = "SELECT externalid, long_url, created_at FROM short_urls"
	InsertOne       = "INSERT INTO short_urls (externalid, long_url, created_at) VALUES($1, $2, $3) RETURNING id"
	GetByExternalID = "SELECT id, externalid, long_url, created_at FROM short_urls where externalId = $1 "
)

func (repo *URLRepository) GetAll() ([]model.URL, error) {
	rows, err := Db.Query(GetAll)
	if err != nil {
		return nil, err
	}
	urls, err := repo.urlFactory.FromSQLRows(rows)
	log.Printf("Returning all URLs")
	return urls, err
}

func (repo *URLRepository) GetByExternalId(externalID string) (model.URL, error) {
	var url model.URL
	row := Db.QueryRow(GetByExternalID, externalID)
	err := row.Scan(&url.ID, &url.ExternalID, &url.LongURL, &url.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return model.URL{}, urlErrors.FromExternalID(externalID)

	}
	if err != nil {
		return model.URL{}, err
	}
	return url, nil
}

func (repo *URLRepository) Add(url model.URL) (model.URL, error) {
	var newID int
	err := Db.QueryRow(InsertOne, url.ExternalID, url.LongURL, time.Now()).Scan(&newID)
	log.Printf("New URL added succesfully with id %d", newID)

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		log.Printf("Error during adding new url with externalId %s, with error: %v", url.ExternalID, err)
		if pqErr.Code == UniqueConstraintViolation { // unique_violation
			return model.URL{}, ErrUniqueConstraint
		}
	}
	return url, err
}
