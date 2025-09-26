package hugocontext

import (
	"fmt"
	"testing"
)

func TestDump_1(t *testing.T) {
	type args struct {
		source []byte
		level  int
	}
	tests := []struct {
		name string
		n    *HugoContext
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

			tt.n.Dump(tt.args.source, tt.args.level)
		})
	}
}
