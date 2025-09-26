package allconfig

import (
	"fmt"
	"testing"
	"time"
)

func TestCloneForLang_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Config{
		C: &ConfigCompiled{
			Timeout: 10 * time.Second,
		},
	}

	x := c.cloneForLang()

	if x.C == nil {
		t.Errorf("x.C was nil")
	}

	if x.C.Timeout != 10*time.Second {
		t.Errorf("x.C.Timeout was %v", x.C.Timeout)
	}
}
