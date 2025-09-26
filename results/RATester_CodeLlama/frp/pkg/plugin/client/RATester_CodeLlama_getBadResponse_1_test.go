package client

import (
	"fmt"
	"testing"
)

func TestGetBadResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	res := getBadResponse()
	if res.Status != "407 Not authorized" {
		t.Errorf("Status = %v, want %v", res.Status, "407 Not authorized")
	}
	if res.StatusCode != 407 {
		t.Errorf("StatusCode = %v, want %v", res.StatusCode, 407)
	}
}
