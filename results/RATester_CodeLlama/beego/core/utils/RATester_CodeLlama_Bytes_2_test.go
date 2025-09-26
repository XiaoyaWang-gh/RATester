package utils

import (
	"fmt"
	"testing"
)

func TestBytes_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Email{
		To: []string{"test@example.com"},
	}
	_, err := e.Bytes()
	if err != nil {
		t.Errorf("Failed to render message headers: %s", err)
	}
}
