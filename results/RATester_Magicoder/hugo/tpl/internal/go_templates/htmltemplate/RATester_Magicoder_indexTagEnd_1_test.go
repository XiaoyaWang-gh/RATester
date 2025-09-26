package template

import (
	"fmt"
	"testing"
)

func TestindexTagEnd_1(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		tag  []byte
		want int
	}{
		{
			name: "Test case 1",
			s:    []byte("<field name=\"string\">"),
			tag:  []byte("name"),
			want: 10,
		},
		{
			name: "Test case 2",
			s:    []byte("<field name=\"string\">"),
			tag:  []byte("want"),
			want: -1,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := indexTagEnd(tt.s, tt.tag); got != tt.want {
				t.Errorf("indexTagEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
