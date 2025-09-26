package gin

import (
	"fmt"
	"testing"
)

func TestAssert1_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Log("recovered", r)
		}
	}()
	assert1(false, "test")
	t.Error("should panic")
}
