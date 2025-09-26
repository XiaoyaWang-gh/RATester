package log

import (
	"fmt"
	"testing"
)

func TestErrorf_1(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestErrorf_0",
			args: args{
				format: "test %s",
				v:      []any{"value"},
			},
		},
		{
			name: "TestErrorf_1",
			args: args{
				format: "test %d",
				v:      []any{123},
			},
		},
		{
			name: "TestErrorf_2",
			args: args{
				format: "test %f",
				v:      []any{123.456},
			},
		},
		{
			name: "TestErrorf_3",
			args: args{
				format: "test %v",
				v:      []any{struct{}{}},
			},
		},
		{
			name: "TestErrorf_4",
			args: args{
				format: "test %+v",
				v:      []any{struct{}{}},
			},
		},
		{
			name: "TestErrorf_5",
			args: args{
				format: "test %#v",
				v:      []any{struct{}{}},
			},
		},
		{
			name: "TestErrorf_6",
			args: args{
				format: "test %T",
				v:      []any{struct{}{}},
			},
		},
		{
			name: "TestErrorf_7",
			args: args{
				format: "test %%",
				v:      []any{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Errorf(tt.args.format, tt.args.v...)
		})
	}
}
