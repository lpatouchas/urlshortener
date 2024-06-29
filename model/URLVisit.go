package model

import (
	"math/big"
	"time"
)

type URLVisit struct {
	ID        big.Int   `json:"-"`
	URL       URL       `json:"url"`
	VisitedAt time.Time `json:"visitedAt"`
}
