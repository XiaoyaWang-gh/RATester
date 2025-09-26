package gin

import (
	"fmt"
	"testing"
)

func TestRecovery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := Recovery()
	ctx := &Context{}

	handler(ctx)

	// Add assertions here
}
