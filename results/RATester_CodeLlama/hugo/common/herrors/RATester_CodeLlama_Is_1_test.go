package herrors

import (
	"fmt"
	"testing"
)

func TestIs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &FeatureNotAvailableError{}
	if !e.Is(e) {
		t.Errorf("Is() = false, want true")
	}
	if e.Is(nil) {
		t.Errorf("Is(nil) = true, want false")
	}
	if e.Is(&FeatureNotAvailableError{}) {
		t.Errorf("Is(&FeatureNotAvailableError{}) = true, want false")
	}
}
