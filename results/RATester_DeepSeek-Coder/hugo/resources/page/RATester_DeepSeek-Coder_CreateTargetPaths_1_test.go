package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateTargetPaths_1(t *testing.T) {
	type args struct {
		d TargetPathDescriptor
	}
	tests := []struct {
		name   string
		args   args
		wantTp TargetPaths
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

			gotTp := CreateTargetPaths(tt.args.d)
			if !reflect.DeepEqual(gotTp, tt.wantTp) {
				t.Errorf("CreateTargetPaths() = %v, want %v", gotTp, tt.wantTp)
			}
		})
	}
}
