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
		path:     "path",
		name:     "name",
		policies: "policies",
	}
	if got := e.Error(); got != "access denied: name is not whitelisted in policy path; the current security configuration is:\n\npolicies\n\n" {
		t.Errorf("AccessDeniedError.Error() = %v, want %v", got, "access denied: name is not whitelisted in policy path; the current security configuration is:\n\npolicies\n\n")
	}
}
