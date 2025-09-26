package v1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewVisitorConfigurerByType_1(t *testing.T) {
	type args struct {
		t VisitorType
	}
	tests := []struct {
		name string
		args args
		want VisitorConfigurer
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

			if got := NewVisitorConfigurerByType(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVisitorConfigurerByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
