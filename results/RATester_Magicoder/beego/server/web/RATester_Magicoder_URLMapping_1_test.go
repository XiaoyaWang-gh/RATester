package web

import (
	"fmt"
	"testing"
)

func TestURLMapping_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{}
	ctrl.URLMapping()
}
