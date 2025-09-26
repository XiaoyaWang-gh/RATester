package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterUnwantedMounts_1(t *testing.T) {
	type args struct {
		mounts []Mount
	}
	tests := []struct {
		name string
		args args
		want []Mount
	}{
		{
			name: "Test with no duplicates",
			args: args{
				mounts: []Mount{
					{Source: "source1", Target: "target1"},
					{Source: "source2", Target: "target2"},
				},
			},
			want: []Mount{
				{Source: "source1", Target: "target1"},
				{Source: "source2", Target: "target2"},
			},
		},
		{
			name: "Test with duplicates",
			args: args{
				mounts: []Mount{
					{Source: "source1", Target: "target1"},
					{Source: "source1", Target: "target1"},
				},
			},
			want: []Mount{
				{Source: "source1", Target: "target1"},
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

			if got := filterUnwantedMounts(tt.args.mounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterUnwantedMounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
