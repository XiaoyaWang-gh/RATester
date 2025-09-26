package binding

import (
	"fmt"
	"testing"
)

func TestName_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := queryBinding{}
	name := qb.Name()
	if name != "query" {
		t.Errorf("Expected 'query', but got '%s'", name)
	}
}
