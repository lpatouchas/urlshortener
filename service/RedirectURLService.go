package service

type RedirectURLService interface {
	GetRedirectURL(urlExternalId string) (string, error)
}
