package highlight

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnormalizeHighlightOptions_1(t *testing.T) {
	tests := []struct {
		name string
		args map[string]any
		want map[string]any
	}{
		{
			name: "Test case 1",
			args: map[string]any{
				"linosStartKey": 1,
				"noHlKey":       true,
				"lineNosKey":    "table",
				"hlLinesKey":    [][2]int{{1, 2}, {3, 4}},
			},
			want: map[string]any{
				"linosstartkey":     1,
				"noHlKey":           true,
				"linenoskey":        "table",
				"hllineskey_parsed": [][2]int{{2, 3}, {4, 5}},
			},
		},
		{
			name: "Test case 2",
			args: map[string]any{
				"linosStartKey": 2,
				"noHlKey":       false,
				"lineNosKey":    "inline",
				"hlLinesKey":    [][2]int{{2, 3}, {4, 5}},
			},
			want: map[string]any{
				"linosstartkey":     2,
				"noHlKey":           false,
				"linenoskey":        "inline",
				"hllineskey_parsed": [][2]int{{3, 4}, {5, 6}},
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

			normalizeHighlightOptions(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("normalizeHighlightOptions() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
