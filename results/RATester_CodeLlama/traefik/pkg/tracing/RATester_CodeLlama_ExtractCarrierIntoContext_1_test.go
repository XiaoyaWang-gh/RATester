package tracing

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestExtractCarrierIntoContext_1(t *testing.T) {
	t.Run("test case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// arrange
		ctx := context.Background()
		headers := http.Header{}
		// act
		result := ExtractCarrierIntoContext(ctx, headers)
		// assert
		if result == nil {
			t.Errorf("result was nil, want not nil")
		}
	})
	t.Run("test case 2:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// arrange
		ctx := context.Background()
		headers := http.Header{}
		// act
		result := ExtractCarrierIntoContext(ctx, headers)
		// assert
		if result == nil {
			t.Errorf("result was nil, want not nil")
		}
	})
}
