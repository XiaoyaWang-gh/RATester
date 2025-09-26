package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRollback_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	db := &dbQueryLog{}
	// when:
	err := db.Rollback()
	// then:
	assert.NoError(t, err)
}
