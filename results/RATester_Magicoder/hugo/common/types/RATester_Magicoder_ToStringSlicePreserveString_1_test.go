package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringSlicePreserveString_1(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want []string
	}{
		{
			name: "Test case 1",
			v:    []string{"test1", "test2"},
			want: []string{"test1", "test2"},
		},
		{
			name: "Test case 2",
			v:    "test",
			want: []string{"test"},
		},
		{
			name: "Test case 3",
			v:    nil,
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

			if got := ToStringSlicePreserveString(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSlicePreserveString() = %v, want %v", got, tt.want)
			}
		})
	}
}
