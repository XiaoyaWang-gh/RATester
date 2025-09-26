package fiber

import (
	"fmt"
	"testing"
)

func TestStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{}
	code := 302
	r.Status(code)

	if r.status != code {
		t.Errorf("Expected status code %d, but got %d", code, r.status)
	}
}
