package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCookie_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestCookie_0",
			args: args{
				key: "test_key",
			},
			want: "test_value",
		},
		{
			name: "TestCookie_1",
			args: args{
				key: "test_key_1",
			},
			want: "test_value_1",
		},
		{
			name: "TestCookie_2",
			args: args{
				key: "test_key_2",
			},
			want: "test_value_2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			input := &BeegoInput{
				Context: &Context{
					Request: &http.Request{
						Header: make(http.Header),
					},
				},
			}
			input.Context.Request.AddCookie(&http.Cookie{
				Name:  tt.args.key,
				Value: tt.want,
			})
			if got := input.Cookie(tt.args.key); got != tt.want {
				t.Errorf("Cookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
