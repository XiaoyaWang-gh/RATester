package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func Testclean_1(t *testing.T) {
	tests := []struct {
		name string
		rm   *RootMapping
		want *RootMapping
	}{
		{
			name: "Test case 1",
			rm: &RootMapping{
				From: "test/",
				To:   "test/",
			},
			want: &RootMapping{
				From: "test",
				To:   "test",
			},
		},
		{
			name: "Test case 2",
			rm: &RootMapping{
				From: "test/../test",
				To:   "test/../test",
			},
			want: &RootMapping{
				From: "test",
				To:   "test",
			},
		},
		{
			name: "Test case 3",
			rm: &RootMapping{
				From: "test/..",
				To:   "test/..",
			},
			want: &RootMapping{
				From: ".",
				To:   ".",
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

			tt.rm.clean()
			if !reflect.DeepEqual(tt.rm, tt.want) {
				t.Errorf("clean() = %v, want %v", tt.rm, tt.want)
			}
		})
	}
}
