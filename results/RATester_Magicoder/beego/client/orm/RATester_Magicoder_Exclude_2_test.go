package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExclude_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{}
	o.Exclude("UserName", "slene")

	// Assuming the Exclude method modifies the querySet's condition field,
	// we can test the condition by checking if it contains the expected condition.
	expectedCondition := NewCondition().AndNot("UserName", "slene")
	if !reflect.DeepEqual(o.cond, expectedCondition) {
		t.Errorf("Expected condition %v, but got %v", expectedCondition, o.cond)
	}
}
