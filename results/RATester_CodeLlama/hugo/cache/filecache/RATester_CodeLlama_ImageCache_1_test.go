package filecache

import (
	"fmt"
	"testing"
)

func TestImageCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := Caches{}
	if got := f.ImageCache(); got != nil {
		t.Errorf("ImageCache() = %v, want nil", got)
	}
}
