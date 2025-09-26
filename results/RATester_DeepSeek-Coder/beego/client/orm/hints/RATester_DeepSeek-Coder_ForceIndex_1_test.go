package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestForceIndex_1(t *testing.T) {
	tests := []struct {
		name    string
		indexes []string
		want    *Hint
	}{
		{
			name:    "Test with one index",
			indexes: []string{"index1"},
			want: &Hint{
				key:   "$forceIndex",
				value: []string{"index1"},
			},
		},
		{
			name:    "Test with multiple indexes",
			indexes: []string{"index1", "index2"},
			want: &Hint{
				key:   "$forceIndex",
				value: []string{"index1", "index2"},
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

			got := ForceIndex(tt.indexes...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForceIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
