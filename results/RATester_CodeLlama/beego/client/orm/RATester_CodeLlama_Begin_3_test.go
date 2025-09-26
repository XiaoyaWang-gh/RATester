package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBegin_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}
	tx, err := d.Begin()
	assert.Nil(t, err)
	assert.NotNil(t, tx)
}
