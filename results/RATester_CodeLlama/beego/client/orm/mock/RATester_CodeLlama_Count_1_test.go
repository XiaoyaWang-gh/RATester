package mock

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQueryM2Mer{}
	count, err := d.Count()
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
