package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcheckRequired_1(t *testing.T) {
	type args struct {
		t   reflect.Type
		src map[string][]string
	}
	tests := []struct {
		name string
		args args
		want MultiError
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

			d := &Decoder{}
			if got := d.checkRequired(tt.args.t, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.checkRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
