package httplib

import (
	"fmt"
	"testing"
)

func TestbuildParamBody_1(t *testing.T) {
	tests := []struct {
		name string
		b    *BeegoHTTPRequest
		want string
	}{
		{
			name: "Test with empty params",
			b: &BeegoHTTPRequest{
				params: map[string][]string{},
			},
			want: "",
		},
		{
			name: "Test with one param",
			b: &BeegoHTTPRequest{
				params: map[string][]string{
					"key1": {"value1"},
				},
			},
			want: "key1=value1",
		},
		{
			name: "Test with multiple params",
			b: &BeegoHTTPRequest{
				params: map[string][]string{
					"key1": {"value1"},
					"key2": {"value2"},
				},
			},
			want: "key1=value1&key2=value2",
		},
		{
			name: "Test with multiple values for one param",
			b: &BeegoHTTPRequest{
				params: map[string][]string{
					"key1": {"value1", "value2"},
				},
			},
			want: "key1=value1&key1=value2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.b.buildParamBody(); got != tt.want {
				t.Errorf("buildParamBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
