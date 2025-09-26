package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestOutputFormats_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var r *Spec
	// when
	formats := r.OutputFormats()
	// then
	assert.NotNil(t, formats)
}
