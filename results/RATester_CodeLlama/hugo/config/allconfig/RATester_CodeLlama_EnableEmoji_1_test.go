package allconfig

import (
	"fmt"
	"testing"
)

func TestEnableEmoji_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.EnableEmoji() {
		t.Error("EnableEmoji() should be false")
	}
}
