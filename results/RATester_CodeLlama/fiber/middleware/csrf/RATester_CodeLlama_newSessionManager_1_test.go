package csrf

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/zeebo/assert"
)

func TestNewSessionManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	s := &session.Store{}
	k := "key"
	// WHEN
	sessionManager := newSessionManager(s, k)
	// THEN
	assert.NotNil(t, sessionManager)
	assert.Equal(t, s, sessionManager.session)
	assert.Equal(t, k, sessionManager.key)
}
