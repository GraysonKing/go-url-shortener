package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "RickAstley"

	// Persist data mapping
	SaveUrlMapping(shortUrl, initialLink, userUUId)

	// Retrieve initial URL
	retrievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retrievedUrl)
}
