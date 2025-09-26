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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := getBehaviorScoresByMode2(tt.args.mode, tt.args.senderScore, tt.args.receiverScore)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBehaviorScoresByMode2() = %v, want %v", got, tt.want)
			}
		})
	}
}
