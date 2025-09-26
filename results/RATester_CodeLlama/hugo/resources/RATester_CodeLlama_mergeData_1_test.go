package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMergeData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &genericResource{}
	in := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}
	// when
	r.mergeData(in)
	// then
	assert.Equal(t, in, r.sd.Data)
}
