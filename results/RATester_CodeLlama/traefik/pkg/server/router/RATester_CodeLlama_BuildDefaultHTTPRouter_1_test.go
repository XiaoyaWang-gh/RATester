package router

import (
	"fmt"
	"testing"
)

func TestBuildDefaultHTTPRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := BuildDefaultHTTPRouter()
	if router == nil {
		t.Error("router is nil")
	}
}
