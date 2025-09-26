package paths

import (
	"fmt"
	"testing"
)

func TestIsLeafBundle_1(t *testing.T) {
	type fields struct {
		bundleType PathType
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test IsLeafBundle when bundleType is PathTypeLeaf",
			fields: fields{
				bundleType: PathTypeLeaf,
			},
			want: true,
		},
		{
			name: "Test IsLeafBundle when bundleType is not PathTypeLeaf",
			fields: fields{
				bundleType: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Path{
				bundleType: tt.fields.bundleType,
			}
			if got := p.IsLeafBundle(); got != tt.want {
				t.Errorf("IsLeafBundle() = %v, want %v", got, tt.want)
			}
		})
	}
}
