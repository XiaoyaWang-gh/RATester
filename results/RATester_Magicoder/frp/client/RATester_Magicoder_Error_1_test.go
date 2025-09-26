package client

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testErr := errors.New("test error")
	ce := cancelErr{Err: testErr}

	if ce.Error() != testErr.Error() {
		t.Errorf("Expected %s, got %s", testErr.Error(), ce.Error())
	}
}
