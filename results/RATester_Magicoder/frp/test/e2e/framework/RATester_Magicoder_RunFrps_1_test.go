package framework

import (
	"fmt"
	"testing"
)

func TestRunFrps_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewFramework(Options{
		TotalParallelNode: 1,
		CurrentNodeIndex:  0,
		FromPortIndex:     10000,
		ToPortIndex:       10001,
	})
	defer f.AfterEach()

	f.BeforeEach()

	_, _, err := f.RunFrps("-c", f.serverConfPaths[0])
	if err != nil {
		t.Fatalf("Failed to start frps: %v", err)
	}
}
