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
	nop := nopConverter(0)
	res, err := nop.Convert(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res == nil {
		t.Error("Expected non-nil result")
	}
}
