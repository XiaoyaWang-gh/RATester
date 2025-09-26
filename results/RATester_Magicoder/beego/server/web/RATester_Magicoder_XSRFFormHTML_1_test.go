package web

import (
	"fmt"
	"testing"
)

func TestXSRFFormHTML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		_xsrfToken: "test_token",
	}

	expected := `<input type="hidden" name="_xsrf" value="test_token" />`
	result := ctrl.XSRFFormHTML()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
