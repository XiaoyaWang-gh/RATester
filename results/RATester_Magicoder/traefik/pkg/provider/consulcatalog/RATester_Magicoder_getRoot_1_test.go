package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetRoot_1(t *testing.T) {
	tests := []struct {
		name string
		c    *connectCert
		want []types.FileOrContent
	}{
		{
			name: "test_1",
			c: &connectCert{
				root: []string{"root1", "root2"},
			},
			want: []types.FileOrContent{"root1", "root2"},
		},
		{
			name: "test_2",
			c: &connectCert{
				root: []string{"root3", "root4"},
			},
			want: []types.FileOrContent{"root3", "root4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.getRoot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
