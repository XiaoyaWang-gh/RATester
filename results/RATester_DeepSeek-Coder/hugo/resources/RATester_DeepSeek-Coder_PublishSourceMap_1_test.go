package resources

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPublishSourceMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &ResourceTransformationCtx{
		OutPath: "/tmp/test",
		OpenResourcePublisher: func(relTargetPath string) (io.WriteCloser, error) {
			return os.OpenFile(relTargetPath, os.O_RDWR|os.O_CREATE, 0666)
		},
	}

	content := "test content"
	err := ctx.PublishSourceMap(content)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	data, err := os.ReadFile(ctx.OutPath + ".map")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(data) != content {
		t.Errorf("Expected content to be '%s', got '%s'", content, string(data))
	}
}
