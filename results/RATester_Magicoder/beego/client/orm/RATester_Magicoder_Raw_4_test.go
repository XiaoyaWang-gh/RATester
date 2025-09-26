package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRaw_4(t *testing.T) {
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name string
		f    *filterOrmDecorator
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

			if got := tt.f.Raw(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOrmDecorator.Raw() = %v, want %v", got, tt.want)
			}
		})
	}
}
