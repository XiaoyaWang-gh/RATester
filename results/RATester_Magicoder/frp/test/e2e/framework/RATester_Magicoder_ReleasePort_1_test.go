package framework

import (
	"fmt"
	"testing"
)

func TestReleasePort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewFramework(Options{
		TotalParallelNode: 1,
		CurrentNodeIndex:  0,
		FromPortIndex:     10000,
		ToPortIndex:       20000,
	})
	defer f.AfterEach()

	port := f.AllocPort()
	f.ReleasePort(port)

	// Add your test case here
}
