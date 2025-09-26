package collections

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestApply_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ns := &Namespace{}

	// Test case 1: Apply function on a nil value
	_, err := ns.Apply(ctx, nil, "apply")
	if err == nil || err.Error() != "can't iterate over a nil value" {
		t.Errorf("Expected error 'can't iterate over a nil value', but got %v", err)
	}

	// Test case 2: Apply function on a non-slice/array value
	_, err = ns.Apply(ctx, 123, "apply")
	if err == nil || err.Error() != "can't apply over 123" {
		t.Errorf("Expected error 'can't apply over 123', but got %v", err)
	}

	// Test case 3: Apply function on a slice
	slice := []int{1, 2, 3, 4, 5}
	result, err := ns.Apply(ctx, slice, "apply")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []any{nil, nil, nil, nil, nil}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 4: Apply function on a slice with a non-existent function
	_, err = ns.Apply(ctx, slice, "nonExistentFunction")
	if err == nil || err.Error() != "can't find function nonExistentFunction" {
		t.Errorf("Expected error 'can't find function nonExistentFunction', but got %v", err)
	}

	// Test case 5: Apply function on a slice with a function that returns an error
	_, err = ns.Apply(ctx, slice, "apply")
	if err == nil || err.Error() != "can't apply myself (no turtles allowed)" {
		t.Errorf("Expected error 'can't apply myself (no turtles allowed)', but got %v", err)
	}
}
