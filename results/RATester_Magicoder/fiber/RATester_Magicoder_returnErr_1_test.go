package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestreturnErr_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		b    *Bind
		args args
		want error
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

			if got := tt.b.returnErr(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
