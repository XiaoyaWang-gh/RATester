package gin

import (
	"fmt"
	"testing"
)

func TestauthorizationHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	user := "testuser"
	password := "testpassword"
	expected := "Basic dGVzdHVzZXJAZGFyYW1zLmNvbTp0ZXN0cGFzc3dvcmQ="
	result := authorizationHeader(user, password)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
