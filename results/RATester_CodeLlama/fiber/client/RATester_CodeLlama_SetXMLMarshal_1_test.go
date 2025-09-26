package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetXMLMarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{}
	f := func(v any) ([]byte, error) {
		return nil, nil
	}
	// when
	c.SetXMLMarshal(f)
	// then
	assert.Equal(t, f, c.XMLMarshal())
}
