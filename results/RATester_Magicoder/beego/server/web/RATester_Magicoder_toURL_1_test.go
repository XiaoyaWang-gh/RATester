package web

import (
	"fmt"
	"testing"
)

func TesttoURL_1(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]string
		want   string
	}{
		{
			name:   "empty params",
			params: map[string]string{},
			want:   "",
		},
		{
			name:   "single param",
			params: map[string]string{"key1": "value1"},
			want:   "key1=value1",
		},
		{
			name:   "multiple params",
			params: map[string]string{"key1": "value1", "key2": "value2"},
			want:   "key1=value1&key2=value2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toURL(tt.params); got != tt.want {
				t.Errorf("toURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
