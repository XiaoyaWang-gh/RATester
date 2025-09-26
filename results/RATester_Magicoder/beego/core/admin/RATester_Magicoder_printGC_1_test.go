package admin

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
)

func TestprintGC_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	memStats := &runtime.MemStats{}
	gcstats := &debug.GCStats{}
	w := &bytes.Buffer{}

	printGC(memStats, gcstats, w)

	// Add assertions here to check the output of the function
}
