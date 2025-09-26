package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testindirect_1(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want any
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

			if got := indirect(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirect() = %v, want %v", got, tt.want)
			}
		})
	}
}
