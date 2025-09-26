package csrf

import (
	"fmt"
	"testing"
)

func TestTokenFromContext_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
