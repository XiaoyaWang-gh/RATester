package deploy

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/spf13/afero"
)

func TestContentType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lf := &localFile{
		NativePath: "test.txt",
		fs:         afero.NewMemMapFs(),
		matcher:    &deployconfig.Matcher{ContentType: "text/plain"},
	}
	if got := lf.ContentType(); got != "text/plain" {
		t.Errorf("ContentType() = %v, want %v", got, "text/plain")
	}
}
