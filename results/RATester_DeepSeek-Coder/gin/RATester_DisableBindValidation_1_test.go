package gin

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin/binding"
)

func TestDisableBindValidation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	DisableBindValidation()
	if binding.Validator != nil {
		t.Errorf("Expected binding.Validator to be nil after calling DisableBindValidation, but got %v", binding.Validator)
	}
}
