package mirror

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeader_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}

	// Test that calling WriteHeader does not panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("blackHoleResponseWriter.WriteHeader() panicked: %v", r)
			}
		}()
		b.WriteHeader(http.StatusOK)
	}()
}
