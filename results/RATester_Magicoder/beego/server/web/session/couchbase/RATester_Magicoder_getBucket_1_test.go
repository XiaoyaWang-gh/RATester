package couchbase

import (
	"fmt"
	"testing"
)

func TestgetBucket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		SavePath: "localhost:8091",
		Pool:     "default",
		Bucket:   "test",
	}

	bucket := provider.getBucket()

	if bucket == nil {
		t.Error("Expected a bucket, but got nil")
	}
}
