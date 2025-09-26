package echo

import (
	"fmt"
	"testing"
)

func TestNewDefaultFS_1(t *testing.T) {
	tests := []struct {
		name string
		want *defaultFS
	}{
		{
			name: "Test case 1",
			want: &defaultFS{
				prefix: "./",
				fs:     nil,
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

			got := newDefaultFS()
			if got.prefix != tt.want.prefix {
				t.Errorf("newDefaultFS() = %v, want %v", got, tt.want)
			}
			if got.fs != tt.want.fs {
				t.Errorf("newDefaultFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
