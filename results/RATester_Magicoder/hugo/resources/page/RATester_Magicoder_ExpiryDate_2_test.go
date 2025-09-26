package page

import (
	"fmt"
	"testing"
)

func TestExpiryDate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	expiryDate := nop.ExpiryDate()

	if !expiryDate.IsZero() {
		t.Errorf("Expected expiry date to be zero, but got: %v", expiryDate)
	}
}
