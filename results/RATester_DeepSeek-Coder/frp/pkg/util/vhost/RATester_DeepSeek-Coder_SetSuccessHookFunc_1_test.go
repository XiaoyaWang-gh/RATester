package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetSuccessHookFunc_1(t *testing.T) {
	type args struct {
		f successHookFunc
	}
	tests := []struct {
		name string
		args args
		want *Muxer
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

			v := &Muxer{}
			if got := v.SetSuccessHookFunc(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Muxer.SetSuccessHookFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
