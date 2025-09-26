package gateway

import (
	"fmt"
	"testing"
)

func TestTranslateNotFoundError_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err error
	if _, err = translateNotFoundError(err); err != nil {
		t.Errorf("translateNotFoundError() error = %v", err)
	}
}
