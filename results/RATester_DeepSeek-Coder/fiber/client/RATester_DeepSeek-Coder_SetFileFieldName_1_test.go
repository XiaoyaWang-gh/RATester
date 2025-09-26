package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFileFieldName_1(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want SetFileFunc
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

			if got := SetFileFieldName(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFileFieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}
