package attributes

import (
	"fmt"
	"testing"
)

func TestDump_2(t *testing.T) {
	type args struct {
		source []byte
		level  int
	}
	tests := []struct {
		name string
		a    *attributesBlock
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

			tt.a.Dump(tt.args.source, tt.args.level)
		})
	}
}
