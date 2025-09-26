package types

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestToStrArray_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Domain{
		Main: "test.com",
		SANs: []string{"test1.com", "test2.com"},
	}
	assert.Equal(t, []string{"test.com", "test1.com", "test2.com"}, d.ToStrArray())
}
