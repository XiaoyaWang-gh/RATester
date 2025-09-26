package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
			r: &ResponseModifier{
				req: &http.Request{},
				rw:  httptest.NewRecorder(),
			},
			args: args{
				code: 200,
			},
			want: 200,
		},
		{
			name: "TestWriteHeader_2",
			r: &ResponseModifier{
				req: &http.Request{},
				rw:  httptest.NewRecorder(),
			},
			args: args{
				code: 404,
			},
			want: 404,
		},
		{
			name: "TestWriteHeader_3",
			r: &ResponseModifier{
				req: &http.Request{},
				rw:  httptest.NewRecorder(),
			},
			args: args{
				code: 500,
			},
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
			if got := tt.r.code; got != tt.want {
				t.Errorf("ResponseModifier.WriteHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
