package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBehaviorScoresByMode2_1(t *testing.T) {
	type args struct {
		mode          int
		senderScore   int
		receiverScore int
	}
	tests := []struct {
		name string
		args args
		want []*BehaviorScore
	}{
		{
			name: "Test Case 1",
			args: args{
				mode:          1,
				senderScore:   5,
				receiverScore: 10,
			},
			want: []*BehaviorScore{
				{Mode: 1, Index: 0, Score: 5},
				{Mode: 1, Index: 1, Score: 10},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				mode:          2,
				senderScore:   0,
				receiverScore: -5,
			},
			want: []*BehaviorScore{
				{Mode: 2, Index: 0, Score: 0},
				{Mode: 2, Index: 1, Score: -5},
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

			if got := getBehaviorScoresByMode2(tt.args.mode, tt.args.senderScore, tt.args.receiverScore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBehaviorScoresByMode2() = %v, want %v", got, tt.want)
			}
		})
	}
}
