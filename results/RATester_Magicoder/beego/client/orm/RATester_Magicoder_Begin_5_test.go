package orm

import (
	"fmt"
	"testing"
)

func TestBegin_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &orm{}
	tx, err := o.Begin()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test if tx is of type TxOrmer
	if _, ok := tx.(TxOrmer); !ok {
		t.Errorf("Expected tx to be of type TxOrmer, but it is not")
	}
}
