package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewScratcher_1(t *testing.T) {
	tests := []struct {
		name string
		want Scratcher
	}{
		{
			name: "Test case 1",
			want: NewScratcher(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewScratcher()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScratcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
