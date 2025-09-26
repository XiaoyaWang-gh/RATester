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

	q := queryBinding{}
	expected := "query"
	actual := q.Name()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
