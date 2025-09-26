package web

import (
	"fmt"
	"testing"
)

func TestFilter_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	action := "before"
	filter := []HandleFunc{}
	n.Filter(action, filter...)
}
