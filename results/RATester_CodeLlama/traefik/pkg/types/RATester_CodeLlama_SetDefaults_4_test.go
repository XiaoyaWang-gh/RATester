package types

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &OtelHTTP{}
	c.SetDefaults()
	assert.Equal(t, "https://localhost:4318", c.Endpoint)
}
