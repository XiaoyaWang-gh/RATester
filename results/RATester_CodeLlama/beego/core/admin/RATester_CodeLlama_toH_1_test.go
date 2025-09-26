package admin

import (
	"fmt"
	"testing"
)

func TestToH_1(t *testing.T) {
	var tests = []struct {
		name  string
		bytes uint64
		want  string
	}{
		{
			name:  "1023",
			bytes: 1023,
			want:  "1023B",
		},
		{
			name:  "1024",
			bytes: 1024,
			want:  "1.00K",
		},
		{
			name:  "1024*1024",
			bytes: 1024 * 1024,
			want:  "1.00M",
		},
		{
			name:  "1024*1024*1024",
			bytes: 1024 * 1024 * 1024,
			want:  "1.00G",
		},
		{
			name:  "1024*1024*1024*1024",
			bytes: 1024 * 1024 * 1024 * 1024,
			want:  "1.00G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toH(tt.bytes); got != tt.want {
				t.Errorf("toH() = %v, want %v", got, tt.want)
			}
		})
	}
}
