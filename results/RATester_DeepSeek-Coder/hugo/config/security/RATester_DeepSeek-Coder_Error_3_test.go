package security

import (
	"fmt"
	"testing"
)

func TestError_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &AccessDeniedError{
		path:     "test_path",
		name:     "test_name",
		policies: "test_policies",
	}

	expected := fmt.Sprintf("access denied: %q is not whitelisted in policy %q; the current security configuration is:\n\n%s\n\n", e.name, e.path, e.policies)
	actual := e.Error()

	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
