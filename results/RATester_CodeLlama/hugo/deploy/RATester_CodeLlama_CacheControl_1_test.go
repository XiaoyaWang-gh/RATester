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
			CacheControl: "max-age=31536000",
		},
	}
	if got := lf.CacheControl(); got != "max-age=31536000" {
		t.Errorf("CacheControl() = %v, want %v", got, "max-age=31536000")
	}
}
