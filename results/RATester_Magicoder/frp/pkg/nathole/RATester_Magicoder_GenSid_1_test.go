package nathole

import (
	"fmt"
	"testing"
)

func TestGenSid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{}
	sid := ctrl.GenSid()
	if sid == "" {
		t.Error("Expected a non-empty string, but got an empty string")
	}
}
