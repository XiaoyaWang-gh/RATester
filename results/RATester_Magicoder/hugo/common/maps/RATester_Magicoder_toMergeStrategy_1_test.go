package maps

import (
	"fmt"
	"testing"
)

func TesttoMergeStrategy_1(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want ParamsMergeStrategy
	}{
		{
			name: "Test case 1",
			v:    "Deep",
			want: ParamsMergeStrategyDeep,
		},
		{
			name: "Test case 2",
			v:    "None",
			want: ParamsMergeStrategyNone,
		},
		{
			name: "Test case 3",
			v:    "Shallow",
			want: ParamsMergeStrategyShallow,
		},
		{
			name: "Test case 4",
			v:    "Invalid",
			want: ParamsMergeStrategyDeep,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toMergeStrategy(tt.v); got != tt.want {
				t.Errorf("toMergeStrategy() = %v, want %v", got, tt.want)
			}
		})
	}
}
