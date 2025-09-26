package paths

import (
	"fmt"
	"testing"
)

func TestSection_1(t *testing.T) {
	tests := []struct {
		name string
		p    *Path
		want string
	}{
		{
			name: "Test case 1",
			p: &Path{
				s:              "test/section",
				posSectionHigh: 5,
			},
			want: "section",
		},
		{
			name: "Test case 2",
			p: &Path{
				s:              "test/",
				posSectionHigh: 0,
			},
			want: "",
		},
		{
			name: "Test case 3",
			p: &Path{
				s:              "test",
				posSectionHigh: 4,
			},
			want: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.Section(); got != tt.want {
				t.Errorf("Section() = %v, want %v", got, tt.want)
			}
		})
	}
}
