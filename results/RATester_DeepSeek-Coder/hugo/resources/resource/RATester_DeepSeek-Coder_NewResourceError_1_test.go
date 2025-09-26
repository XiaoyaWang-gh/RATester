package resource

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewResourceError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	testErr := errors.New("test error")
	testData := "test data"

	re := NewResourceError(testErr, testData)

	if re.Error() != testErr.Error() {
		t.Errorf("Expected error message to be '%s', got '%s'", testErr.Error(), re.Error())
	}

	if re.Data() != testData {
		t.Errorf("Expected data to be '%v', got '%v'", testData, re.Data())
	}
}
