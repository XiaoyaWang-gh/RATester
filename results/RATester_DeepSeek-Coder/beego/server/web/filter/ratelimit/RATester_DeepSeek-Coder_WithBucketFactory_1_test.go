package ratelimit

import (
	"fmt"
	"testing"
)

func TestWithBucketFactory_1(t *testing.T) {
	testCases := []struct {
		name string
		opt  limiterOption
	}{
		{
			name: "WithBucketFactory",
			opt:  WithBucketFactory(func(opts ...bucketOption) bucket { return nil }),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &limiter{}
			tc.opt(l)
			if l.bucketFactory == nil {
				t.Errorf("Expected bucketFactory to be set, but it was not")
			}
		})
	}
}
