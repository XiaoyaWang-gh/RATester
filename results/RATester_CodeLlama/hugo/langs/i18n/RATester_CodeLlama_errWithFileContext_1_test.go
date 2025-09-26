package i18n

import (
	"fmt"
	"testing"
)

func TestErrWithFileContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add test cases.
}
