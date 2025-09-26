package capture

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestFromContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	capt := &Capture{}
	ctx := context.WithValue(context.Background(), capturedData, capt)

	captured, err := FromContext(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(captured, *capt) {
		t.Errorf("expected %v, got %v", *capt, captured)
	}
}
