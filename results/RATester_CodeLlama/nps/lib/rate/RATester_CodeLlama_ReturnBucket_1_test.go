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
		bucketSize: 100,
	}
	rate.ReturnBucket(100)
	if rate.NowRate != 100 {
		t.Errorf("rate.NowRate = %d, want %d", rate.NowRate, 100)
	}
}
