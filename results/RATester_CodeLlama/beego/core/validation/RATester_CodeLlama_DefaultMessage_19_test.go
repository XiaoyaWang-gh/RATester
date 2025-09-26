package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Required{}
	if r.DefaultMessage() != MessageTmpls["Required"] {
		t.Errorf("DefaultMessage() = %v, want %v", r.DefaultMessage(), MessageTmpls["Required"])
	}
}
