package images

import (
	"fmt"
	"testing"
)

func TestHeight_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &Image{}
	if got := i.Height(); got != 0 {
		t.Errorf("Height() = %v, want %v", got, 0)
	}
}
