package deploy

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestReader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lf := &localFile{
		NativePath: "/path/to/file",
		matcher:    &deployconfig.Matcher{Gzip: true},
		gzipped:    *bytes.NewBuffer([]byte("gzipped contents")),
	}

	reader, err := lf.Reader()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(content) != "gzipped contents" {
		t.Errorf("Expected 'gzipped contents', got '%s'", string(content))
	}
}
