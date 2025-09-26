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

	err := errors.New("test error")
	data := "test data"
	resourceError := NewResourceError(err, data)
	if resourceError.Error() != err.Error() {
		t.Errorf("NewResourceError() error = %v, want %v", resourceError.Error(), err.Error())
	}
	if resourceError.Data() != data {
		t.Errorf("NewResourceError() data = %v, want %v", resourceError.Data(), data)
	}
}
