package session

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRefresh_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &Session{
		id: "1234567890",
	}
	// when
	s.refresh()
	// then
	assert.NotEqual(t, "1234567890", s.id)
}
