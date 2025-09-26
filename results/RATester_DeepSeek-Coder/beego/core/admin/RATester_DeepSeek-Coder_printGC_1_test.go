package admin

import (
	"fmt"
	"io"
	"runtime"
	"runtime/debug"
	"testing"
)

func TestPrintGC_1(t *testing.T) {
	type args struct {
		memStats *runtime.MemStats
		gcstats  *debug.GCStats
		w        io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			printGC(tt.args.memStats, tt.args.gcstats, tt.args.w)
		})
	}
}
