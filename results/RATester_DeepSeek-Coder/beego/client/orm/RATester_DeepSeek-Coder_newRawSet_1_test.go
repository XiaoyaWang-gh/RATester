package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRawSet_1(t *testing.T) {
	type args struct {
		orm   *ormBase
		query string
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want RawSeter
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

			if got := newRawSet(tt.args.orm, tt.args.query, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRawSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
