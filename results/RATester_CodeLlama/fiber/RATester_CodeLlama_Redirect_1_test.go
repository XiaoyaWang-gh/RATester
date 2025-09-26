package fiber

import (
	"fmt"
	"testing"
)

func TestRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.Redirect()
}
