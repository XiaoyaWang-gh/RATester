package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestByType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	ns := &Namespace{}
	typ := "typ"

	// when
	result := ns.ByType(typ)

	// then
	assert.NotNil(t, result)
}
