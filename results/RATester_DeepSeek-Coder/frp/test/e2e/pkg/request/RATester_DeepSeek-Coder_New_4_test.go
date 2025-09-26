package request

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_4(t *testing.T) {
	tests := []struct {
		name string
		want *Request
	}{
		{
			name: "TestNew",
			want: &Request{
				protocol: "tcp",
				addr:     "127.0.0.1",
				method:   "GET",
				path:     "/",
				headers:  map[string]string{},
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

			got := New()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
