package deploy

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestOpenBucket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &Deployer{
		target: &deployconfig.Target{
			Name: "test",
			URL:  "https://example.com",
		},
	}
	bucket, err := d.openBucket(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if bucket == nil {
		t.Fatal("expected bucket")
	}
}
