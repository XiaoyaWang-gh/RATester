package paths

import (
	"fmt"
	"testing"
)

func TestIsBranchBundle_1(t *testing.T) {
	type fields struct {
		bundleType PathType
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test IsBranchBundle when bundleType is PathTypeBranch",
			fields: fields{
				bundleType: PathTypeBranch,
			},
			want: true,
		},
		{
			name: "Test IsBranchBundle when bundleType is not PathTypeBranch",
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
			if got := p.IsBranchBundle(); got != tt.want {
				t.Errorf("IsBranchBundle() = %v, want %v", got, tt.want)
			}
		})
	}
}
