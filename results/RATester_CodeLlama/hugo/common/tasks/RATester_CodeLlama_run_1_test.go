package tasks

import (
	"fmt"
	"testing"
)

func TestRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{}
	r.run()
}
