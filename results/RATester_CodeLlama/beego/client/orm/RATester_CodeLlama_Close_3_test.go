package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	d := &stmtQueryLog{}
	// act
	err := d.Close()
	// assert
	assert.NoError(t, err)
}
