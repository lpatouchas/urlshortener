package model

import (
	"time"
)

type URL struct {
	ID         int       `json:"-"`
	ExternalID string    `json:"externalId"`
	LongURL    string    `json:"longUrl"`
	CreatedAt  time.Time `json:"createdAt" swagger:"str fmt date-time"` //TODO still as sting on swagger
}

type NewURL struct {
	LongURL string `json:"longUrl" binding:"required,url"`
}
