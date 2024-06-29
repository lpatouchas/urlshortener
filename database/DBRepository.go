package database

import "urlshortener/model"

type DBRepository[T model.GenericModel] interface {
	GetAll() ([]T, error)
	Add(entity T) (T, error)
	GetByExternalId(externalID string) (model.URL, error)
}
