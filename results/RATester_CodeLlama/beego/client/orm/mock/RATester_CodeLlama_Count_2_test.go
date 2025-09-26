package mock

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestCount_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	count, err := d.Count()
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}
