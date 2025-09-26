package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWrapSite_1(t *testing.T) {
	type args struct {
		s Site
	}
	tests := []struct {
		name string
		args args
		want Site
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

			if got := WrapSite(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
