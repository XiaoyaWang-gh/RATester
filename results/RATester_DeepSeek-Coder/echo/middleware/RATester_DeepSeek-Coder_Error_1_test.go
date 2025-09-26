package middleware

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

	testErr := &ErrKeyAuthMissing{Err: errors.New("test error")}
	expected := "test error"

	if testErr.Error() != expected {
		t.Errorf("Expected %s, got %s", expected, testErr.Error())
	}
}
