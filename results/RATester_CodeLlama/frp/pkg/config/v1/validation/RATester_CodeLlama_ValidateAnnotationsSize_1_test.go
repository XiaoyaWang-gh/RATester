package validation

import (
	"fmt"
	"testing"
)

func TestValidateAnnotationsSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var annotations = map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	err := ValidateAnnotationsSize(annotations)
	if err != nil {
		t.Errorf("ValidateAnnotationsSize failed, err: %v", err)
	}
}
