package context

import (
	"fmt"
	"testing"
)

func TestError_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := StatusCode(404)
	if s.Error() != "404" {
		t.Errorf("Error() = %v, want %v", s.Error(), "404")
	}
}
