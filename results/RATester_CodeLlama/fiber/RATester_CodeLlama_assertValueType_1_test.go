package fiber

import (
	"fmt"
	"testing"
)

func TestAssertValueType_1(t *testing.T) {
	t.Parallel()
	t.Run("Test assertValueType", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		result := assertValueType[int, int](1)
		if result != 1 {
			t.Errorf("assertValueType() = %v, want %v", result, 1)
		}
	})
}
