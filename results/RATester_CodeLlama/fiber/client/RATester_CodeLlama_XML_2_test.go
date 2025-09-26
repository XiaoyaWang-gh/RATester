package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestXML_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &Response{}
	v := &struct{}{}
	// when
	err := r.XML(v)
	// then
	assert.NoError(t, err)
}
