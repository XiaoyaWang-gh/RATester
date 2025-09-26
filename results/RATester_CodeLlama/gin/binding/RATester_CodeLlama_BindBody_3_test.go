package binding

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
	"github.com/stretchr/testify/require"
)

func TestBindBody_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	body := []byte{0x81, 0x88, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}
	var obj string
	// when
	err := msgpackBinding{}.BindBody(body, &obj)
	// then
	require.NoError(t, err)
	assert.Equal(t, "hello world", obj)
}
