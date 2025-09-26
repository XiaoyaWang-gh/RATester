package deploy

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestInvalidateCloudFront_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	target := &deployconfig.Target{
		URL:                      "https://example.com",
		CloudFrontDistributionID: "E123456789",
	}

	err := InvalidateCloudFront(ctx, target)
	if err != nil {
		t.Errorf("InvalidateCloudFront() error = %v", err)
	}
}
