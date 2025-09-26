package page

import (
	"fmt"
	"testing"
)

func TestServerPort_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	s := testSite{}
	if s.ServerPort() != 1313 {
		t.Errorf("ServerPort() = %v, want %v", s.ServerPort(), 1313)
	}
}
