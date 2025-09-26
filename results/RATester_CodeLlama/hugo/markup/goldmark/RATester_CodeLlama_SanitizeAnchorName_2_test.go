package goldmark

import (
	"fmt"
	"testing"
)

func TestSanitizeAnchorName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &goldmarkConverter{}
	s := "test"
	expected := "test"
	actual := c.SanitizeAnchorName(s)
	if actual != expected {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
