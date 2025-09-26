package types

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Domain{}
	domains := []string{"test.com", "test2.com"}
	d.Set(domains)
	assert.Equal(t, "test.com", d.Main)
	assert.Equal(t, []string{"test2.com"}, d.SANs)
}
