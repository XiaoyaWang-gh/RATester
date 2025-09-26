package gin

import (
	"fmt"
	"testing"
)

func TestAssert1_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			if r != "fail" {
				t.Errorf("Expected panic message to be 'fail', got '%v'", r)
			}
		}
	}()

	assert1(false, "fail")
}
