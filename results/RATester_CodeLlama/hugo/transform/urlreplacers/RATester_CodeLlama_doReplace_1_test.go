package urlreplacers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/transform"
)

func TestDoReplace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "path"
	ct := transform.FromTo(nil)
	quotes := [][]byte{[]byte("\""), []byte("\"")}

	doReplace(path, ct, quotes)

	// TODO: add assertions.
}
