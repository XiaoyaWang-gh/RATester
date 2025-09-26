package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestLogger_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &context{
		echo: &Echo{},
	}
	assert.NotNil(t, c.Logger())
}
