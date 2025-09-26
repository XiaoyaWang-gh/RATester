package helpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestReaderToBytes_1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		lines := strings.NewReader("hello world")

		// when
		result := ReaderToBytes(lines)

		// then
		assert.Equal(t, []byte("hello world"), result)
	})
}
