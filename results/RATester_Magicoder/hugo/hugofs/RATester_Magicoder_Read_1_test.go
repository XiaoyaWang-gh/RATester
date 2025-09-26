package hugofs

import (
	"fmt"
	"testing"
)

func TestRead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &noOpRegularFileOps{}
	_, err := file.Read([]byte{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
