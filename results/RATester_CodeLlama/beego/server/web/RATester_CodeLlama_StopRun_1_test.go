package web

import (
	"fmt"
	"testing"
)

func TestStopRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.StopRun()
}
