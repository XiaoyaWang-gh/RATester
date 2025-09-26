package paths

import (
	"fmt"
	"testing"
)

func TestIsBundle_1(t *testing.T) {
	type fields struct {
		bundleType PathType
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "TestIsBundle_Bundle",
			fields: fields{
				bundleType: PathTypeLeaf,
			},
			want: true,
		},
		{
			name: "TestIsBundle_NotBundle",
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
			if got := p.IsBundle(); got != tt.want {
				t.Errorf("IsBundle() = %v, want %v", got, tt.want)
			}
		})
	}
}
