package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetInternal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	he := &HTTPError{}
	err := errors.New("test error")
	he.SetInternal(err)
	if he.Internal != err {
		t.Errorf("Expected Internal to be %v, but got %v", err, he.Internal)
	}
}
