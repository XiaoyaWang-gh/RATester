package tls

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Certificates{}
	assert.Equal(t, "certificates", c.Type())
}
