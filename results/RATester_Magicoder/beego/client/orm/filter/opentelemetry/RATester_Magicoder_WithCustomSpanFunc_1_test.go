package opentelemetry

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithCustomSpanFunc_1(t *testing.T) {
	type args struct {
		customSpanFunc CustomSpanFunc
	}
	tests := []struct {
		name string
		args args
		want FilterChainOption
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

			if got := WithCustomSpanFunc(tt.args.customSpanFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCustomSpanFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
