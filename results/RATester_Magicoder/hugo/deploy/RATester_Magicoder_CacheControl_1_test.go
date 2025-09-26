package deploy

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestCacheControl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lf := &localFile{
		matcher: &deployconfig.Matcher{
			CacheControl: "public, max-age=31536000",
		},
	}

	expected := "public, max-age=31536000"
	result := lf.CacheControl()

	if result != expected {
		t.Errorf("Expected CacheControl to be '%s', but got '%s'", expected, result)
	}
}
