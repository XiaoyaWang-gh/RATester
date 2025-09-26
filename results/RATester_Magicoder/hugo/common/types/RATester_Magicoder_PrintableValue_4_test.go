package types

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestPrintableValue_4(t *testing.T) {
	type args struct {
		r Result[int]
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test case 1",
			args: args{
				r: Result[int]{
					Value: 10,
					Err:   nil,
				},
			},
			want: 10,
		},
		{
			name: "Test case 2",
			args: args{
				r: Result[int]{
					Value: 0,
					Err:   errors.New("test error"),
				},
			},
			want: errors.New("test error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.r.PrintableValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintableValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
