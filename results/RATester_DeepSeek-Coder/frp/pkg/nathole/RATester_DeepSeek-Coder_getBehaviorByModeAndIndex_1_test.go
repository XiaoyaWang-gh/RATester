package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBehaviorByModeAndIndex_1(t *testing.T) {
	type args struct {
		mode  int
		index int
	}
	tests := []struct {
		name  string
		args  args
		wantA RecommandBehavior
		wantB RecommandBehavior
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

			gotA, gotB := getBehaviorByModeAndIndex(tt.args.mode, tt.args.index)
			if !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("getBehaviorByModeAndIndex() gotA = %v, want %v", gotA, tt.wantA)
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("getBehaviorByModeAndIndex() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
