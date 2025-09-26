package nathole

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/samber/lo"
)

func TestGetBehaviorByMode_1(t *testing.T) {
	type args struct {
		mode int
	}
	tests := []struct {
		name string
		args args
		want []lo.Tuple2[RecommandBehavior, RecommandBehavior]
	}{
		{
			name: "Test case 1",
			args: args{mode: 0},
			want: mode0Behaviors,
		},
		{
			name: "Test case 2",
			args: args{mode: 1},
			want: mode1Behaviors,
		},
		{
			name: "Test case 3",
			args: args{mode: 2},
			want: mode2Behaviors,
		},
		{
			name: "Test case 4",
			args: args{mode: 3},
			want: mode3Behaviors,
		},
		{
			name: "Test case 5",
			args: args{mode: 4},
			want: mode4Behaviors,
		},
		{
			name: "Test case 6",
			args: args{mode: 5},
			want: mode0Behaviors,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getBehaviorByMode(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBehaviorByMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
