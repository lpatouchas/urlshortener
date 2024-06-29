package factory

import (
	"time"
	"urlshortener/model"
)

type URLVisitFactory struct {
}

//func (URLVisitFactory *URLVisitFactory) FromSQLRows(rows *sql.Rows) ([]model.URLVisit, error) {
//	defer rows.Close()
//
//	var shortURLs []model.URLVisit
//	for rows.Next() {
//		var urlVisit model.URLVisit
//		if err := rows.Scan(&urlVisit.URL, &urlVisit.LongURL, &shortURL.CreatedAt); err != nil {
//			return nil, err
//		}
//		shortURLs = append(shortURLs, shortURL)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return shortURLs, nil
//}

func (URLVisitFactory *URLVisitFactory) FromURL(url model.URL) model.URLVisit {
	return model.URLVisit{URL: url, VisitedAt: time.Now()}
}
