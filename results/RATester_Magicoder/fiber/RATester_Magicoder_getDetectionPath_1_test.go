package fiber

import (
	"fmt"
	"testing"
)

func TestgetDetectionPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		detectionPath: "testPath",
	}

	got := ctx.getDetectionPath()
	want := "testPath"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
