package memory

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	s := &Storage{
		db: map[string]entry{
			"key": {
				data: []byte("value"),
			},
		},
	}
	key := "key"

	// act
	value, err := s.Get(key)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, []byte("value"), value)
}
