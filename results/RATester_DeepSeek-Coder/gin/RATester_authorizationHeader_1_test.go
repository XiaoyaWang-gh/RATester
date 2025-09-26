package gin

import (
	"fmt"
	"testing"
)

func TestAuthorizationHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	user := "testUser"
	password := "testPassword"
	expected := "Basic dGVzdHVzZXI6dGVzdHBhc3N3b3Jk"
	result := authorizationHeader(user, password)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
