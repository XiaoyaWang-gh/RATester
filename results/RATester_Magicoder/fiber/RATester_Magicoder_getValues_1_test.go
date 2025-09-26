package fiber

import (
	"fmt"
	"testing"
)

func TestgetValues_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		values: [maxParams]string{
			"value1",
			"value2",
			"value3",
			// ...
		},
	}

	result := ctx.getValues()

	if result != &ctx.values {
		t.Errorf("Expected %p, got %p", &ctx.values, result)
	}
}
