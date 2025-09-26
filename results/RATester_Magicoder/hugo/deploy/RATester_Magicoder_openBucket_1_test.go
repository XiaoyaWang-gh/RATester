package deploy

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestopenBucket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &Deployer{
		target: &deployconfig.Target{
			Name: "test",
			URL:  "test://test",
		},
	}
	bucket, err := d.openBucket(ctx)
	if err != nil {
		t.Fatalf("openBucket failed: %v", err)
	}
	if bucket == nil {
		t.Fatal("openBucket returned nil bucket")
	}
}
