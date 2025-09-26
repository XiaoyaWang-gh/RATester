package framework

import (
	"errors"
	"fmt"
	"testing"
)

func TestExpectNoError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	err := errors.New("test error")
	ExpectNoError(err)
}
