package i18n

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToPluralCountValue_1(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want any
	}{
		{
			name: "Test with float",
			in:   1.2,
			want: "1.2",
		},
		{
			name: "Test with int",
			in:   1,
			want: 1,
		},
		{
			name: "Test with string",
			in:   "1.2",
			want: "1.2",
		},
		{
			name: "Test with string that can be converted to float",
			in:   "1",
			want: "1.0",
		},
		{
			name: "Test with unsupported type",
			in:   []string{"1", "2"},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toPluralCountValue(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPluralCountValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
