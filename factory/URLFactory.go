package factory

import (
	"database/sql"
	"time"
	"urlshortener/model"
)

type URLFactory struct {
}

func (URLFactory *URLFactory) FromSQLRows(rows *sql.Rows) ([]model.URL, error) {
	defer rows.Close()

	var shortURLs []model.URL
	for rows.Next() {
		var shortURL model.URL
		if err := rows.Scan(&shortURL.ExternalID, &shortURL.LongURL, &shortURL.CreatedAt); err != nil {
			return nil, err
		}
		shortURLs = append(shortURLs, shortURL)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shortURLs, nil
}

func (URLFactory *URLFactory) FromNewURL(newURL model.NewURL) model.URL {
	externalId := GenerateRandomString()
	return model.URL{ExternalID: externalId, LongURL: newURL.LongURL, CreatedAt: time.Now()}
}
