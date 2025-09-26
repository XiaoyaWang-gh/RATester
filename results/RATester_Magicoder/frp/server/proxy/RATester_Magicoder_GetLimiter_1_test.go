package proxy

import (
	"fmt"
	"testing"

	"golang.org/x/time/rate"
)

func TestGetLimiter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		limiter: rate.NewLimiter(1, 1),
	}

	limiter := pxy.GetLimiter()

	if limiter != pxy.limiter {
		t.Errorf("Expected limiter to be %v, but got %v", pxy.limiter, limiter)
	}
}
