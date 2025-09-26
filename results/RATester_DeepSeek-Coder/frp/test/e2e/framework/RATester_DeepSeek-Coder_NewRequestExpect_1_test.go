package framework

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/test/e2e/framework/consts"
)

func TestNewRequestExpect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testFramework := &Framework{
		TempDirectory: "/tmp",
	}

	reqExpect := NewRequestExpect(testFramework)

	if reqExpect.f != testFramework {
		t.Errorf("Expected framework to be %v, got %v", testFramework, reqExpect.f)
	}

	if string(reqExpect.expectResp) != consts.TestString {
		t.Errorf("Expected response to be %s, got %s", consts.TestString, string(reqExpect.expectResp))
	}

	if reqExpect.expectError != false {
		t.Errorf("Expected error to be false, got %v", reqExpect.expectError)
	}
}
