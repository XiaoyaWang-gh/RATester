package client

import (
	"fmt"
	"testing"
)

func TestPathParam_1(t *testing.T) {
	r := &Request{
		path: &PathParam{
			"key": "value",
		},
	}

	tests := []struct {
		name string
		key  string
		want string
	}{
		{
			name: "Test case 1",
			key:  "key",
			want: "value",
		},
		{
			name: "Test case 2",
			key:  "non-existing-key",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := r.PathParam(tt.key); got != tt.want {
				t.Errorf("PathParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
