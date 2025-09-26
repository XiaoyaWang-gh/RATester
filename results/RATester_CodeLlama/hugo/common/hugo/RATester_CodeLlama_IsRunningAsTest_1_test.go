package hugo

import (
	"fmt"
	"os"
	"testing"
)

func TestIsRunningAsTest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if IsRunningAsTest() {
		t.Error("IsRunningAsTest() should return false when running as test")
	}
	os.Args = append(os.Args, "-test.v")
	if !IsRunningAsTest() {
		t.Error("IsRunningAsTest() should return true when running as test")
	}
}
