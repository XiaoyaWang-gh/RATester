package nathole

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/samber/lo"
)

func TestGetBehaviorByMode_1(t *testing.T) {
	tests := []struct {
		name string
		mode int
		want []lo.Tuple2[RecommandBehavior, RecommandBehavior]
	}{
		{
			name: "test mode 0",
			mode: 0,
			want: mode0Behaviors,
		},
		{
			name: "test mode 1",
			mode: 1,
			want: mode1Behaviors,
		},
		{
			name: "test mode 2",
			mode: 2,
			want: mode2Behaviors,
		},
		{
			name: "test mode 3",
			mode: 3,
			want: mode3Behaviors,
		},
		{
			name: "test mode 4",
			mode: 4,
			want: mode4Behaviors,
		},
		{
			name: "test mode 5",
			mode: 5,
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

			if got := getBehaviorByMode(tt.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBehaviorByMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
