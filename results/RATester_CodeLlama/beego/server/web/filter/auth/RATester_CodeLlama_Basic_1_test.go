package auth

import (
	"fmt"
	"testing"
)

func TestBasic_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	username := "admin"
	password := "123456"
	filter := Basic(username, password)
	if filter == nil {
		t.Errorf("Basic() = %v, want %v", filter, nil)
	}
}
