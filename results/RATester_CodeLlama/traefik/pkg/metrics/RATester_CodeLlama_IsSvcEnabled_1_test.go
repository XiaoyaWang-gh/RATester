package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsSvcEnabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &standardRegistry{}
	r.svcEnabled = true
	assert.True(t, r.IsSvcEnabled())
}
