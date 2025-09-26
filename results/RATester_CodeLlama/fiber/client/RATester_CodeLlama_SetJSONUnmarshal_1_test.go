package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetJSONUnmarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{}
	f := func(data []byte, v any) error {
		return nil
	}
	// when
	c.SetJSONUnmarshal(f)
	// then
	assert.Equal(t, f, c.jsonUnmarshal)
}
