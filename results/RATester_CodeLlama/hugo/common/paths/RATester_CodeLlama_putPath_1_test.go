package paths

import (
	"fmt"
	"testing"
)

func TestPutPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	p := &Path{}
	// When
	putPath(p)
	// Then
	if got, want := pathPool.Get(), p; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
