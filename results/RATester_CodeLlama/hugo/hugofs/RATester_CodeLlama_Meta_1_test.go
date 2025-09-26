package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMeta_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	fi := &dirEntryMeta{}
	// when
	actual := fi.Meta()
	// then
	assert.NotNil(t, actual)
}
