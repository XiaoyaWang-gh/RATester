package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Request: &http.Request{},
	}
	if c.Done() != nil {
		t.Errorf("Done() = %v, want nil", c.Done())
	}
}
