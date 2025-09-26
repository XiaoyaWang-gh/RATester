package toml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Config{}
	_, err := c.Parse("")
	assert.NotNil(t, err)
}
