package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestC_1(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{
			name: "Test C",
			want: &Client{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := C(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("C() = %v, want %v", got, tt.want)
			}
		})
	}
}
