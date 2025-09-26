package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestDefaultExpiredFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expiredFunc := defaultExpiredFunc()

	for i := 0; i < 10; i++ {
		expired := expiredFunc()
		if expired < time.Duration(0) {
			t.Errorf("Expected a positive duration, got %v", expired)
		}
		if expired > (time.Duration(5) * time.Second) {
			t.Errorf("Expected a duration less than or equal to 5 seconds, got %v", expired)
		}
	}
}
