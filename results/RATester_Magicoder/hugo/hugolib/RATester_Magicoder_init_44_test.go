package hugolib

import (
	"fmt"
	"testing"
)

func Testinit_44(t *testing.T) {
	type args struct {
		preserveOringal bool
	}
	tests := []struct {
		name string
		m    *pageMetaParams
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

			tt.m.init(tt.args.preserveOringal)
		})
	}
}
