package pagesfromdata

import (
	"fmt"
	"testing"
)

func Testhash_3(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		b    *BuildState
		args args
		want uint64
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

			if got := tt.b.hash(tt.args.v); got != tt.want {
				t.Errorf("BuildState.hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
