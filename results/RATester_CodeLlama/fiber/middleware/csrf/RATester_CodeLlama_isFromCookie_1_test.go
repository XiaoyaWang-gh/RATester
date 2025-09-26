package csrf

import (
	"fmt"
	"testing"
)

func TestIsFromCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	extractor := FromCookie
	if !isFromCookie(extractor) {
		t.Errorf("isFromCookie() = %v, want %v", isFromCookie(extractor), true)
	}
}
