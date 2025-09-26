package context

import (
	"fmt"
	"testing"
)

func TestProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{}
	input.Header("X-Forwarded-For")
	input.Proxy()
}
