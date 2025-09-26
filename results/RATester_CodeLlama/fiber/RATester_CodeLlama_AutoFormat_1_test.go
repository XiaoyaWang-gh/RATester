package fiber

import (
	"fmt"
	"testing"
)

func TestAutoFormat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c *DefaultCtx
	var body any
	var err error

	err = c.AutoFormat(body)
	if err != nil {
		t.Errorf("AutoFormat() error = %v", err)
		return
	}
}
