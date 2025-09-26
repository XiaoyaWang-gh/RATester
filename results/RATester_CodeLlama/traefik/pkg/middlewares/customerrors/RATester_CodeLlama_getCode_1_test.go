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

	cc := &codeCatcher{}
	cc.code = 200
	if cc.getCode() != 200 {
		t.Errorf("Expected 200, got %d", cc.getCode())
	}
}
