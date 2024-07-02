package factory

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestExternalIDFactory_GenerateRandomString(t *testing.T) {
	// Seed the random number generator for reproducibility in tests.
	err := godotenv.Load("../tests/.env")
	if err != nil {
		println("error loading test envs")
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))

	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "check length of generated string",
			test: func(t *testing.T) {
				generatedString := GenerateRandomString()
				assert.Equal(t, ExternalIDLength, len(generatedString))
			},
		},
		{
			name: "check character set of generated string",
			test: func(t *testing.T) {
				generatedString := GenerateRandomString()
				for _, char := range generatedString {
					assert.Contains(t, ExternalIDCharset, string(char))
				}
			},
		},
		{
			name: "check randomness of generated string",
			test: func(t *testing.T) {
				generatedStrings := make(map[string]bool)
				for i := 0; i < 10000; i++ {
					generatedString := GenerateRandomString()
					// Check if generated string is already in the map (non-randomness check)
					if _, exists := generatedStrings[generatedString]; exists {
						t.Fatalf("generated string is not random: %s", generatedString)
					}
					generatedStrings[generatedString] = true
				}
				assert.Greater(t, len(generatedStrings), 1, "randomness check failed")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}
