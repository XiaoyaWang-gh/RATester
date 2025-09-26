package rate

import (
	"fmt"
	"io"
	"testing"

	"gotest.tools/assert"
)

func TestRead_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	conn := &rateConn{}
	conn.rate = &Rate{}
	// when
	n, err := conn.Read([]byte{})
	// then
	assert.Equal(t, 0, n)
	assert.Equal(t, io.EOF, err)
}
