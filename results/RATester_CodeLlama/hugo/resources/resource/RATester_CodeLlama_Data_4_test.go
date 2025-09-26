package resource

import (
	"errors"
	"fmt"
	"testing"
)

func TestData_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &resourceError{
		error: errors.New("error"),
		data:  "data",
	}
	if e.Data() != "data" {
		t.Errorf("Data() = %v, want %v", e.Data(), "data")
	}
}
