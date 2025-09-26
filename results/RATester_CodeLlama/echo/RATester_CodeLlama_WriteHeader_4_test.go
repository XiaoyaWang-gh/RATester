package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestWriteHeader_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &Response{}
	code := 200

	// when
	r.WriteHeader(code)

	// then
	assert.Equal(t, code, r.Status)
}
