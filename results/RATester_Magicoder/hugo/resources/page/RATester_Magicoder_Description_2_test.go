package page

import (
	"fmt"
	"testing"
)

func TestDescription_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	description := nop.Description()
	if description != "" {
		t.Errorf("Expected empty string, got %s", description)
	}
}
