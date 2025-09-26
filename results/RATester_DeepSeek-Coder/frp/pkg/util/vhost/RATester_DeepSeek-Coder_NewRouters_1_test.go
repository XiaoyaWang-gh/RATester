package vhost

import (
	"fmt"
	"testing"
)

func TestNewRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := NewRouters()
	if r == nil {
		t.Errorf("NewRouters() = %v, want not nil", r)
	}
	if r.indexByDomain == nil {
		t.Errorf("NewRouters().indexByDomain = %v, want not nil", r.indexByDomain)
	}
}
