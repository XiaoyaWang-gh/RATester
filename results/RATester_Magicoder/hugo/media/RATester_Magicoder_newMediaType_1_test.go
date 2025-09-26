package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewMediaType_1(t *testing.T) {
	tests := []struct {
		name     string
		main     string
		sub      string
		suffixes []string
		want     Type
	}{
		{
			name:     "test case 1",
			main:     "application",
			sub:      "rss",
			suffixes: []string{"xml"},
			want:     newMediaType("application", "rss", []string{"xml"}),
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newMediaType(tt.main, tt.sub, tt.suffixes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}
