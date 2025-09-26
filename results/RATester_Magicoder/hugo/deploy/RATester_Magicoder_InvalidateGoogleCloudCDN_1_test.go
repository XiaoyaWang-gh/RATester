package deploy

import (
	"context"
	"fmt"
	"testing"
)

func TestInvalidateGoogleCloudCDN_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	origin := "project/origin"

	// Test case 1: Valid origin
	err := InvalidateGoogleCloudCDN(ctx, origin)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Invalid origin
	err = InvalidateGoogleCloudCDN(ctx, "invalid")
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 3: Empty origin
	err = InvalidateGoogleCloudCDN(ctx, "")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
