package session

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/internal/storage/memory"
	"github.com/zeebo/assert"
)

func TestNew_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{
		{
			Storage: memory.New(),
		},
	}

	// Act
	store := New(config...)

	// Assert
	assert.NotNil(t, store)
	assert.NotNil(t, store.Config)
	assert.NotNil(t, store.Config.Storage)
	assert.NotNil(t, store.Config.KeyGenerator)
	assert.NotNil(t, store.Config.KeyLookup)
	assert.NotNil(t, store.Config.CookieDomain)
	assert.NotNil(t, store.Config.CookiePath)
	assert.NotNil(t, store.Config.CookieSameSite)
	assert.NotNil(t, store.Config.source)
	assert.NotNil(t, store.Config.sessionName)
	assert.NotNil(t, store.Config.Expiration)
	assert.NotNil(t, store.Config.CookieSecure)
	assert.NotNil(t, store.Config.CookieHTTPOnly)
	assert.NotNil(t, store.Config.CookieSessionOnly)
}
