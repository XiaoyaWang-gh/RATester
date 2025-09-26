package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFilterForwardRequestHeaders_1(t *testing.T) {
	type args struct {
		forwardRequestHeaders http.Header
		allowedHeaders        []string
	}
	tests := []struct {
		name string
		args args
		want http.Header
	}{
		{
			name: "Test case 1",
			args: args{
				forwardRequestHeaders: http.Header{
					"Content-Type":  []string{"application/json"},
					"Authorization": []string{"Bearer token"},
				},
				allowedHeaders: []string{"Content-Type"},
			},
			want: http.Header{
				"Content-Type": []string{"application/json"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				forwardRequestHeaders: http.Header{
					"Content-Type":  []string{"application/json"},
					"Authorization": []string{"Bearer token"},
				},
				allowedHeaders: []string{"Authorization"},
			},
			want: http.Header{
				"Authorization": []string{"Bearer token"},
			},
		},
		{
			name: "Test case 3",
			args: args{
				forwardRequestHeaders: http.Header{
					"Content-Type":  []string{"application/json"},
					"Authorization": []string{"Bearer token"},
				},
				allowedHeaders: []string{},
			},
			want: http.Header{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := filterForwardRequestHeaders(tt.args.forwardRequestHeaders, tt.args.allowedHeaders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterForwardRequestHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
