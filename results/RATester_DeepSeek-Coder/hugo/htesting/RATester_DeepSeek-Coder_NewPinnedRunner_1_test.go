package htesting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPinnedRunner_1(t *testing.T) {
	type args struct {
		t            testing.TB
		pinnedTestRe string
	}
	tests := []struct {
		name string
		args args
		want *PinnedRunner
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

			if got := NewPinnedRunner(tt.args.t, tt.args.pinnedTestRe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPinnedRunner() = %v, want %v", got, tt.want)
			}
		})
	}
}
