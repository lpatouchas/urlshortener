package service

import "urlshortener/model"

type GetByExternalID interface {
	GetByExternalID(externalID string) (model.URL, error)
}
