package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestToLabels_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lvs := labelNamesValues{"foo", "bar", "baz", "qux"}
	labels := lvs.ToLabels()
	assert.Equal(t, map[string]string{"foo": "bar", "baz": "qux"}, labels)
}
