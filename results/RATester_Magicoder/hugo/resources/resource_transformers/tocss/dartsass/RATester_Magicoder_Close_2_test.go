package dartsass

import (
	"fmt"
	"testing"

	godartsassv1 "github.com/bep/godartsass"
)

func TestClose_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		transpilerV1: &godartsassv1.Transpiler{},
	}

	err := client.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
