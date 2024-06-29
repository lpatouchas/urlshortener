package factory

import (
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//const externalIDLength = 6

var externalIDLength int
var charset string

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	err := godotenv.Load()
	length, err := strconv.Atoi(os.Getenv("URL_EXTERNAL_ID_SIZE"))
	if err != nil {
		panic(err)
	}
	externalIDLength = length
	charset = os.Getenv("URL_EXTERNAL_ID_CHARSET")
}

type ExternalIDFactory struct {
}

func (externalIDService *ExternalIDFactory) GenerateRandomString() string {
	result := make([]byte, externalIDLength)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
