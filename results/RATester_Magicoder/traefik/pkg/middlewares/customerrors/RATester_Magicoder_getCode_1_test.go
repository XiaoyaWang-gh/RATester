package customerrors

import (
	"fmt"
	"testing"
)

func TestGetCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &codeCatcher{
		code: 200,
	}

	code := cc.getCode()

	if code != 200 {
		t.Errorf("Expected code to be 200, but got %d", code)
	}
}
