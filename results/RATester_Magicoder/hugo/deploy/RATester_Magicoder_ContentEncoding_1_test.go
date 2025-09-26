package deploy

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestContentEncoding_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lf := &localFile{
		matcher: &deployconfig.Matcher{
			Gzip: true,
		},
	}

	expected := "gzip"
	result := lf.ContentEncoding()

	if result != expected {
		t.Errorf("Expected ContentEncoding to be %s, but got %s", expected, result)
	}
}
