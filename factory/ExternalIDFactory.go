package factory

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

//TODO need to find a way to read .env here through tests as well.
//const externalIDLength = 6

var ExternalIDLength int
var ExternalIDCharset string
var loadOnce sync.Once

type ExternalIDFactory struct {
}

func loadEnv() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	length, err := strconv.Atoi(os.Getenv("URL_EXTERNAL_ID_SIZE"))
	if err != nil {
		panic(err)
	}
	ExternalIDLength = length
	ExternalIDCharset = os.Getenv("URL_EXTERNAL_ID_CHARSET")
}

func GenerateRandomString() string {
	loadOnce.Do(loadEnv)

	result := make([]byte, ExternalIDLength)
	for i := range result {
		result[i] = ExternalIDCharset[rand.Intn(len(ExternalIDCharset))]
	}
	return string(result)
}
