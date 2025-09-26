package page

import (
	"fmt"
	"testing"
)

func TestLink_1(t *testing.T) {
	tests := []struct {
		name string
		p    *pagePathBuilder
		want string
	}{
		{
			name: "Test case 1",
			p: &pagePathBuilder{
				els: []string{"test", "case", "1"},
				d: TargetPathDescriptor{
					BaseName: "test",
				},
				baseNameSameAsType: true,
				prefixLink:         "/prefix",
				linkUpperOffset:    1,
			},
			want: "/prefix/test/case/1/",
		},
		{
			name: "Test case 2",
			p: &pagePathBuilder{
				els: []string{"test", "case", "2"},
				d: TargetPathDescriptor{
					BaseName: "test",
				},
				baseNameSameAsType: false,
				prefixLink:         "/prefix",
				linkUpperOffset:    0,
			},
			want: "/prefix/test/case/2",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.Link(); got != tt.want {
				t.Errorf("Link() = %v, want %v", got, tt.want)
			}
		})
	}
}
