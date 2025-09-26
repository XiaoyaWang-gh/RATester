package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFlash_1(t *testing.T) {
	tests := []struct {
		name string
		want *FlashData
	}{
		{
			name: "Test NewFlash",
			want: &FlashData{
				Data: make(map[string]string),
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

			got := NewFlash()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
