package page

import (
	"fmt"
	"testing"
)

func TestPublishDate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	publishDate := nop.PublishDate()

	if !publishDate.IsZero() {
		t.Errorf("Expected PublishDate to be zero, got %v", publishDate)
	}
}
