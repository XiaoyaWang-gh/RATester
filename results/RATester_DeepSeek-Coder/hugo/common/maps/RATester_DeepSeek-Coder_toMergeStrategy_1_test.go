package maps

import (
	"fmt"
	"testing"
)

func TestToMergeStrategy_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    any
		want ParamsMergeStrategy
	}{
		{
			name: "Test with valid ParamsMergeStrategyDeep",
			v:    ParamsMergeStrategyDeep,
			want: ParamsMergeStrategyDeep,
		},
		{
			name: "Test with valid ParamsMergeStrategyNone",
			v:    "none",
			want: ParamsMergeStrategyNone,
		},
		{
			name: "Test with valid ParamsMergeStrategyShallow",
			v:    "shallow",
			want: "shallow",
		},
		{
			name: "Test with invalid string",
			v:    "invalid",
			want: ParamsMergeStrategyDeep,
		},
		{
			name: "Test with invalid int",
			v:    123,
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
