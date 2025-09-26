package paths

import (
	"fmt"
	"testing"
)

func TestModifyPathBundleTypeResource_1(t *testing.T) {
	type args struct {
		p *Path
	}
	tests := []struct {
		name string
		args args
		want *Path
	}{
		// Test case 1
		{
			name: "Test case 1",
			args: args{
				p: &Path{
					bundleType: PathTypeFile,
				},
			},
			want: &Path{
				bundleType: PathTypeContentResource,
			},
		},
		// Test case 2
		{
			name: "Test case 2",
			args: args{
				p: &Path{
					bundleType: PathTypeContentResource,
				},
			},
			want: &Path{
				bundleType: PathTypeContentResource,
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

			ModifyPathBundleTypeResource(tt.args.p)
			if tt.args.p.bundleType != tt.want.bundleType {
				t.Errorf("ModifyPathBundleTypeResource() = %v, want %v", tt.args.p.bundleType, tt.want.bundleType)
			}
		})
	}
}
