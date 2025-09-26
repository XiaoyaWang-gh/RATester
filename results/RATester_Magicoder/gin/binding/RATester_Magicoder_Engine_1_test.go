package binding

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestEngine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &defaultValidator{}
	expected := &validator.Validate{}
	v.validate = expected
	result := v.Engine()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
