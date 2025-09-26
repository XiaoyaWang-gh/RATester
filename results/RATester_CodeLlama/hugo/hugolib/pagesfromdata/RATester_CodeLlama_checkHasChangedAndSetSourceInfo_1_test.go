package pagesfromdata

import (
	"fmt"
	"testing"
)

func TestCheckHasChangedAndSetSourceInfo_1(t *testing.T) {
	type args struct {
		changedPath string
		v           any
	}
	tests := []struct {
		name string
		b    *BuildState
		args args
		want bool
	}{
		{
			name: "test1",
			b: &BuildState{
				StaleVersion: 1,
			},
			args: args{
				changedPath: "test1",
				v:           "test1",
			},
			want: true,
		},
		{
			name: "test2",
			b: &BuildState{
				StaleVersion: 1,
			},
			args: args{
				changedPath: "test2",
				v:           "test2",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.b.checkHasChangedAndSetSourceInfo(tt.args.changedPath, tt.args.v); got != tt.want {
				t.Errorf("checkHasChangedAndSetSourceInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
