package output

import (
	"fmt"
	"testing"
)

func TestLen_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formats := Formats{
		{Name: "html"},
		{Name: "rss"},
	}

	if formats.Len() != 2 {
		t.Errorf("Expected formats.Len() to be 2, but was %d", formats.Len())
	}
}
