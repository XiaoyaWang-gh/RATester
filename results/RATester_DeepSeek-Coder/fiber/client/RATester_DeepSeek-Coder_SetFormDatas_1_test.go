package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetFormDatas_1(t *testing.T) {
	type args struct {
		m map[string]string
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

			if got := tt.r.SetFormDatas(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.SetFormDatas() = %v, want %v", got, tt.want)
			}
		})
	}
}
