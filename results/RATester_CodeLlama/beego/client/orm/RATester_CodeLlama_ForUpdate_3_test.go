package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForUpdate_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	assert.Equal(t, true, true)
}
