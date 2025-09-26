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

	t.Parallel()
	ctx := context.Background()
	origin := "test-project/test-origin"
	err := InvalidateGoogleCloudCDN(ctx, origin)
	if err != nil {
		t.Errorf("InvalidateGoogleCloudCDN() error = %v", err)
		return
	}
}
