package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNameNormalized_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{
		sd: ResourceSourceDescriptor{
			NameNormalized: "foo",
		},
	}

	assert.Equal(t, "foo", l.NameNormalized())
}
