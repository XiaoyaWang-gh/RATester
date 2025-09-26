package rate

import (
	"fmt"
	"testing"
)

func TestReturnBucket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rate := &Rate{
		bucketSize:        100,
		bucketSurplusSize: 100,
		bucketAddSize:     10,
		stopChan:          make(chan bool),
	}

	rate.ReturnBucket(50)

	if rate.bucketSurplusSize != 150 {
		t.Errorf("Expected bucketSurplusSize to be 150, but got %d", rate.bucketSurplusSize)
	}
}
