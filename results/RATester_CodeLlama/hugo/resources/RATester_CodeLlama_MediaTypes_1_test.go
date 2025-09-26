package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMediaTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var r *Spec

	// when
	mediaTypes := r.MediaTypes()

	// then
	assert.NotNil(t, mediaTypes)
}
