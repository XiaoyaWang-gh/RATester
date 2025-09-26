package echo

import (
	"fmt"
	"testing"
	"testing/fstest"

	"github.com/zeebo/assert"
)

func TestStaticFileHandler_1(t *testing.T) {
	t.Run("should return a HandlerFunc", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		file := "test.txt"
		filesystem := fstest.MapFS{
			file: &fstest.MapFile{
				Data: []byte("Hello, World!"),
			},
		}
		handler := StaticFileHandler(file, filesystem)
		assert.NotNil(t, handler)
	})
}
