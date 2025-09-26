package testing

import (
	"fmt"
	"testing"
)

func TestSetTestingPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testPort := "8080"
	SetTestingPort(testPort)
	if port != testPort {
		t.Errorf("Expected port to be %s, but got %s", testPort, port)
	}
}
