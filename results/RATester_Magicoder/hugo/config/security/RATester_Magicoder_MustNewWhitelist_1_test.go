package security

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustNewWhitelist_1(t *testing.T) {
	type args struct {
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want Whitelist
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

			if got := MustNewWhitelist(tt.args.patterns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustNewWhitelist() = %v, want %v", got, tt.want)
			}
		})
	}
}
