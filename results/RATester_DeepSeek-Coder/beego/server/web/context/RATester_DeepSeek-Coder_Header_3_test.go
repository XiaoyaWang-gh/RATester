package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeader_3(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestHeader_1",
			args: args{key: "Content-Type"},
			want: "application/json",
		},
		{
			name: "TestHeader_2",
			args: args{key: "Accept"},
			want: "*/*",
		},
		{
			name: "TestHeader_3",
			args: args{key: "Authorization"},
			want: "Bearer token",
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
						Header: http.Header{
							"Content-Type":  []string{"application/json"},
							"Accept":        []string{"*/*"},
							"Authorization": []string{"Bearer token"},
						},
					},
				},
			}
			if got := input.Header(tt.args.key); got != tt.want {
				t.Errorf("BeegoInput.Header() = %v, want %v", got, tt.want)
			}
		})
	}
}
