package encryptcookie

import (
	"fmt"
	"testing"
)

func TestNew_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
