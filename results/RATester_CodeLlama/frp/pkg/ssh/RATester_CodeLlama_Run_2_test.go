package ssh

import (
	"fmt"
	"testing"
)

func TestRun_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &Gateway{
		bindPort: 10000,
	}
	g.Run()
}
