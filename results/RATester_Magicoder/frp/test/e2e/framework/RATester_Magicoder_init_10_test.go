package framework

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestInit_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	RunID = uuid.NewString()
	if RunID == "" {
		t.Error("RunID is empty")
	}
}
