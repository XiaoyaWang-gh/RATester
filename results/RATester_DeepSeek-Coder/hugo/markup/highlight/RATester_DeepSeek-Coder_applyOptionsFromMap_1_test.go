package highlight

import (
	"fmt"
	"testing"
)

func TestApplyOptionsFromMap_1(t *testing.T) {
	type args struct {
		optsm map[string]any
		cfg   *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				optsm: map[string]any{
					"Style":              "monokai",
					"CodeFences":         true,
					"NoClasses":          false,
					"NoHl":               false,
					"LineNos":            true,
					"LineNumbersInTable": false,
					"AnchorLineNos":      true,
					"LineAnchors":        "line-anchors",
					"LineNoStart":        1,
					"Hl_Lines":           "1-2 3-4",
					"Hl_inline":          false,
					"TabWidth":           4,
					"GuessSyntax":        true,
				},
				cfg: &Config{},
			},
			wantErr: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := applyOptionsFromMap(tt.args.optsm, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("applyOptionsFromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
