package vhost

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
)

func TestRead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	conn := readOnlyConn{reader: bytes.NewBufferString("hello")}
	p := make([]byte, 5)

	// when
	n, err := conn.Read(p)

	// then
	require.NoError(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "hello", string(p))
}
