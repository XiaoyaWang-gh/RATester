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

	addSize := int64(10)
	r := NewRate(addSize)

	if r.bucketSize != addSize*2 {
		t.Errorf("Expected bucketSize to be %d, but got %d", addSize*2, r.bucketSize)
	}

	if r.bucketSurplusSize != 0 {
		t.Errorf("Expected bucketSurplusSize to be 0, but got %d", r.bucketSurplusSize)
	}

	if r.bucketAddSize != addSize {
		t.Errorf("Expected bucketAddSize to be %d, but got %d", addSize, r.bucketAddSize)
	}

	if r.stopChan == nil {
		t.Error("Expected stopChan to be initialized, but it is nil")
	}
}
