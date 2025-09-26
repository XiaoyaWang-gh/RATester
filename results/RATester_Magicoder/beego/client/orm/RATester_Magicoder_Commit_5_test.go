package orm

import (
	"fmt"
	"testing"
)

func TestCommit_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mockFilter := &filterOrmDecorator{
		insideTx: true,
	}

	err := mockFilter.Commit()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
