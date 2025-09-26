package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestfilterUnwantedMounts_1(t *testing.T) {
	type args struct {
		mounts []Mount
	}
	tests := []struct {
		name string
		args args
		want []Mount
	}{
		{
			name: "Test case 1",
			args: args{
				mounts: []Mount{
					{
						Source: "source1",
						Target: "target1",
					},
					{
						Source: "source2",
						Target: "target2",
					},
					{
						Source: "source1",
						Target: "target1",
					},
				},
			},
			want: []Mount{
				{
					Source: "source1",
					Target: "target1",
				},
				{
					Source: "source2",
					Target: "target2",
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				mounts: []Mount{
					{
						Source: "source1",
						Target: "target1",
					},
					{
						Source: "source1",
						Target: "target1",
					},
					{
						Source: "source2",
						Target: "target2",
					},
				},
			},
			want: []Mount{
				{
					Source: "source1",
					Target: "target1",
				},
				{
					Source: "source2",
					Target: "target2",
				},
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
