package rate

import (
	"fmt"
	"testing"
)

func TestAdd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Rate{
		bucketSize: 10,
	}
	s.add(10)
	if s.bucketSurplusSize != 10 {
		t.Errorf("bucketSurplusSize is %d, want %d", s.bucketSurplusSize, 10)
	}
}
