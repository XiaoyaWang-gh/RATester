package hints

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUseIndex_1(t *testing.T) {
	tests := []struct {
		name    string
		indexes []string
		want    *Hint
	}{
		{
			name:    "Test case 1",
			indexes: []string{"index1", "index2"},
			want: &Hint{
				key:   "UseIndex",
				value: []string{"index1", "index2"},
			},
		},
		{
			name:    "Test case 2",
			indexes: []string{"index3", "index4"},
			want: &Hint{
				key:   "UseIndex",
				value: []string{"index3", "index4"},
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

			got := UseIndex(tt.indexes...)
			if got.GetKey() != tt.want.GetKey() {
				t.Errorf("UseIndex() = %v, want %v", got.GetKey(), tt.want.GetKey())
			}
			if !reflect.DeepEqual(got.GetValue(), tt.want.GetValue()) {
				t.Errorf("UseIndex() = %v, want %v", got.GetValue(), tt.want.GetValue())
			}
		})
	}
}
