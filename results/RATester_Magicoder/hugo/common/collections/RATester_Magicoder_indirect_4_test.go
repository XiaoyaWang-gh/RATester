package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func Testindirect_4(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name      string
		args      args
		wantRv    reflect.Value
		wantIsNil bool
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

			gotRv, gotIsNil := indirect(tt.args.v)
			if !reflect.DeepEqual(gotRv, tt.wantRv) {
				t.Errorf("indirect() gotRv = %v, want %v", gotRv, tt.wantRv)
			}
			if gotIsNil != tt.wantIsNil {
				t.Errorf("indirect() gotIsNil = %v, want %v", gotIsNil, tt.wantIsNil)
			}
		})
	}
}
