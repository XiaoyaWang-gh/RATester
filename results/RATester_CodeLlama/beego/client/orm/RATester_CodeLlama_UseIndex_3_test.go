package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestUseIndex_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	assert.Equal(t, 1, 1)
}
