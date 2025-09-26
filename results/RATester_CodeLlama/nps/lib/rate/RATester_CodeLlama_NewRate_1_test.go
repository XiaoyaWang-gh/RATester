package rate

import (
	"fmt"
	"testing"
)

func TestNewRate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rate := NewRate(10)
	if rate.bucketSize != 20 {
		t.Errorf("rate.bucketSize = %d, want %d", rate.bucketSize, 20)
	}
	if rate.bucketSurplusSize != 0 {
		t.Errorf("rate.bucketSurplusSize = %d, want %d", rate.bucketSurplusSize, 0)
	}
	if rate.bucketAddSize != 10 {
		t.Errorf("rate.bucketAddSize = %d, want %d", rate.bucketAddSize, 10)
	}
	if rate.stopChan == nil {
		t.Errorf("rate.stopChan = %v, want %v", rate.stopChan, nil)
	}
}
