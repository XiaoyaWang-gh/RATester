package herrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestMust_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Must(errors.New("test error"))
}
