package proxy

import (
	"fmt"
	"testing"
)

func TestClose_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	https := &HttpsServer{}
	err := https.Close()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
