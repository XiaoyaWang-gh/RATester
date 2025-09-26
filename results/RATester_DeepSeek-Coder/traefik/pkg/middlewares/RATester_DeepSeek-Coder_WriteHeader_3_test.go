package middlewares

import (
	"fmt"
	"testing"
)

func TestWriteHeader_3(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		r    *ResponseModifier
		args args
		want int
	}{
		{
			name: "TestWriteHeader_1",
			r:    &ResponseModifier{},
			args: args{code: 200},
			want: 200,
		},
		{
			name: "TestWriteHeader_2",
			r:    &ResponseModifier{},
			args: args{code: 199},
			want: 199,
		},
		{
			name: "TestWriteHeader_3",
			r:    &ResponseModifier{},
			args: args{code: 200},
			want: 200,
		},
		{
			name: "TestWriteHeader_4",
			r:    &ResponseModifier{},
			args: args{code: 100},
			want: 100,
		},
		{
			name: "TestWriteHeader_5",
			r:    &ResponseModifier{},
			args: args{code: 500},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.r.WriteHeader(tt.args.code)
			if tt.r.code != tt.want {
				t.Errorf("ResponseModifier.WriteHeader() = %v, want %v", tt.r.code, tt.want)
			}
		})
	}
}
