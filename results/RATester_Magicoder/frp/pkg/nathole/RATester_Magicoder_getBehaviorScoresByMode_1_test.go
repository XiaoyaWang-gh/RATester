package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBehaviorScoresByMode_1(t *testing.T) {
	type args struct {
		mode         int
		defaultScore int
	}
	tests := []struct {
		name string
		args args
		want []*BehaviorScore
	}{
		{
			name: "Test Case 1",
			args: args{
				mode:         1,
				defaultScore: 5,
			},
			want: []*BehaviorScore{
				{
					Mode:  1,
					Index: 0,
					Score: 5,
				},
				{
					Mode:  1,
					Index: 1,
					Score: 5,
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				mode:         2,
				defaultScore: -5,
			},
			want: []*BehaviorScore{
				{
					Mode:  2,
					Index: 0,
					Score: -5,
				},
				{
					Mode:  2,
					Index: 1,
					Score: -5,
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

			if got := getBehaviorScoresByMode(tt.args.mode, tt.args.defaultScore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBehaviorScoresByMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
