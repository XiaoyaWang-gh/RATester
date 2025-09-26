package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := &responseWriter{}

	// Test case 1: Setting a new status code
	rw.WriteHeader(http.StatusOK)
	if rw.status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rw.status)
	}

	// Test case 2: Setting an existing status code
	rw.WriteHeader(http.StatusOK)
	if rw.status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rw.status)
	}

	// Test case 3: Setting a new status code after headers have been written
	rw.WriteHeader(http.StatusInternalServerError)
	if rw.status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rw.status)
	}
}
