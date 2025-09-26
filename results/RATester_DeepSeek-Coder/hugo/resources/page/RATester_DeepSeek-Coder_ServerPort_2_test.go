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

	site := testSite{}
	port := site.ServerPort()
	if port != 1313 {
		t.Errorf("Expected port 1313, got %d", port)
	}
}
