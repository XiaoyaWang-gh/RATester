package fiber

import (
	"fmt"
	"testing"
)

func TestReleaseRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{
		c: &DefaultCtx{
			app: &App{},
		},
	}
	ReleaseRedirect(r)
}
