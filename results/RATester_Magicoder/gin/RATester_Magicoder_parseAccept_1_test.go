package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseAccept_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{
			name: "test case 1",
			arg:  "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
			want: []string{"text/html", "application/xhtml+xml", "application/xml", "*/*"},
		},
		{
			name: "test case 2",
			arg:  "text/html;q=0.9,application/xhtml+xml,application/xml;q=0.8,*/*;q=0.7",
			want: []string{"text/html", "application/xhtml+xml", "application/xml", "*/*"},
		},
		{
			name: "test case 3",
			arg:  "text/html;q=0.9,application/xhtml+xml;q=0.8,application/xml;q=0.7,*/*;q=0.6",
			want: []string{"text/html", "application/xhtml+xml", "application/xml", "*/*"},
		},
		{
			name: "test case 4",
			arg:  "text/html;q=0.9,application/xhtml+xml;q=0.8,application/xml;q=0.7,*/*;q=0.6",
			want: []string{"text/html", "application/xhtml+xml", "application/xml", "*/*"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := parseAccept(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAccept() = %v, want %v", got, tt.want)
			}
		})
	}
}
