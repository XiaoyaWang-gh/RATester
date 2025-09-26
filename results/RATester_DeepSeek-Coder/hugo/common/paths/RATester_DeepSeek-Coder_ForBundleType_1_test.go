package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestForBundleType_1(t *testing.T) {
	type args struct {
		t PathType
	}
	tests := []struct {
		name string
		p    Path
		args args
		want *Path
	}{
		{
			name: "Test ForBundleType",
			p:    Path{},
			args: args{
				t: PathType(1),
			},
			want: &Path{
				bundleType: PathType(1),
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

			if got := tt.p.ForBundleType(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Path.ForBundleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
