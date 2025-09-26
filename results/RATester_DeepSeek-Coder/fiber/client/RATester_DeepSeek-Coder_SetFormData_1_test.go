package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFormData_1(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name string
		r    *Request
		args args
		want *Request
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

			if got := tt.r.SetFormData(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetFormData() = %v, want %v", got, tt.want)
			}
		})
	}
}
