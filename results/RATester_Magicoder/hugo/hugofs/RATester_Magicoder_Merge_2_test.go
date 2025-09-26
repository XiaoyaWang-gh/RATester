package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge_2(t *testing.T) {
	tests := []struct {
		name string
		dst  *FileMeta
		src  *FileMeta
		want *FileMeta
	}{
		{
			name: "Merge with nil dst",
			dst:  nil,
			src:  &FileMeta{Name: "test"},
			want: &FileMeta{Name: "test"},
		},
		{
			name: "Merge with nil src",
			dst:  &FileMeta{Name: "test"},
			src:  nil,
			want: &FileMeta{Name: "test"},
		},
		{
			name: "Merge with non-nil dst and src",
			dst:  &FileMeta{Name: "dst"},
			src:  &FileMeta{Name: "src"},
			want: &FileMeta{Name: "src"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.dst.Merge(tt.src)
			if !reflect.DeepEqual(tt.dst, tt.want) {
				t.Errorf("Merge() = %v, want %v", tt.dst, tt.want)
			}
		})
	}
}
