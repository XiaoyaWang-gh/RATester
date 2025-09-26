package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBind_13(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
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

			if got := Bind(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bind() = %v, want %v", got, tt.want)
			}
		})
	}
}
