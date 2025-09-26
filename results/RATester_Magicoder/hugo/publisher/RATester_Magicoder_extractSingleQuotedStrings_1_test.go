package publisher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestextractSingleQuotedStrings_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "Test case 1",
			s:    "field name 'John Doe'",
			want: []string{"John", "Doe"},
		},
		{
			name: "Test case 2",
			s:    "field name 'John' 'Doe'",
			want: []string{"John", "Doe"},
		},
		{
			name: "Test case 3",
			s:    "field name 'John Doe' 'Jane Doe'",
			want: []string{"John", "Doe", "Jane", "Doe"},
		},
		{
			name: "Test case 4",
			s:    "field name 'John' 'Doe' 'Jane' 'Smith'",
			want: []string{"John", "Doe", "Jane", "Smith"},
		},
		{
			name: "Test case 5",
			s:    "field name 'John' 'Doe' 'Jane' 'Smith'",
			want: []string{"John", "Doe", "Jane", "Smith"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := extractSingleQuotedStrings(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractSingleQuotedStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
