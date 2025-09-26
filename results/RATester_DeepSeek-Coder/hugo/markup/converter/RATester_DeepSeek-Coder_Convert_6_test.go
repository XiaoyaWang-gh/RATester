package converter

import (
	"context"
	"fmt"
	"testing"
)

func TestConvert_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := RenderContext{
		Ctx: context.Background(),
		Src: []byte("test"),
	}

	converter := nopConverter(0)
	result, err := converter.Convert(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Errorf("Expected a non-nil result")
	}

	bytesResult := result.Bytes()
	if len(bytesResult) != 0 {
		t.Errorf("Expected an empty byte slice, got %v", bytesResult)
	}
}
