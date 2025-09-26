package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenamePath_1(t *testing.T) {
	type args struct {
		parentKeyPath string
		m             map[string]any
	}
	tests := []struct {
		name string
		r    KeyRenamer
		args args
		want map[string]any
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

			r := KeyRenamer{
				renames: tt.r.renames,
			}
			r.renamePath(tt.args.parentKeyPath, tt.args.m)
			if !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("renamePath() = %v, want %v", tt.args.m, tt.want)
			}
		})
	}
}
