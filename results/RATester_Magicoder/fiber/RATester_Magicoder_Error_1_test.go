package fiber

import (
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Error{Message: "Test Error"}
	expected := "Test Error"
	if e.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, e.Error())
	}
}
