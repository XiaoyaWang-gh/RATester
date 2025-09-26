package postpub

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestData_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{
		prefix: "prefix",
	}
	m := r.Data()
	assert.Equal(t, map[string]any{
		"Integrity": "",
	}, m)
}
