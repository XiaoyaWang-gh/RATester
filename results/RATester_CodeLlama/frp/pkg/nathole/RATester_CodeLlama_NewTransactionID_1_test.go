package nathole

import (
	"fmt"
	"testing"
)

func TestNewTransactionID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id := NewTransactionID()
	if len(id) != 19 {
		t.Errorf("NewTransactionID() = %v, want %v", len(id), 19)
	}
}
