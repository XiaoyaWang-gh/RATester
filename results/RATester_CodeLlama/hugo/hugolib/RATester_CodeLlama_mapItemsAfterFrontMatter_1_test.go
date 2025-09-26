package hugolib

import (
	"fmt"
	"testing"
)

func TestMapItemsAfterFrontMatter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var rn contentParseInfo
	var s shortcodeHandler
	var source []byte

	// when
	err := rn.mapItemsAfterFrontMatter(source, &s)

	// then
	if err != nil {
		t.Errorf("mapItemsAfterFrontMatter() error = %v", err)
		return
	}
}
