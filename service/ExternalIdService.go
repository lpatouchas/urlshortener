package service

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
const externalIDLength = 6

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

type ExternalIDService struct {
}

func (externalIDService *ExternalIDService) GenerateRandomString() string {
	result := make([]byte, externalIDLength)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
