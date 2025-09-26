package pageparser

import (
	"fmt"
	"testing"
)

func Test__2(t *testing.T) {
	t.Run("tError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if tError != 0 {
			t.Errorf("Expected tError to be 0, got %d", tError)
		}
	})

	t.Run("tEOF", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if tEOF != 1 {
			t.Errorf("Expected tEOF to be 1, got %d", tEOF)
		}
	})

	t.Run("TypeLeadSummaryDivider", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if TypeLeadSummaryDivider != 2 {
			t.Errorf("Expected TypeLeadSummaryDivider to be 2, got %d", TypeLeadSummaryDivider)
		}
	})

	// Continue with the rest of the constants...
}
