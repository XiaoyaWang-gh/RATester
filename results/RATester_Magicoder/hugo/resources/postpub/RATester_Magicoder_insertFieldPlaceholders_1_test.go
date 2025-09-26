package postpub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestinsertFieldPlaceholders_1(t *testing.T) {
	type args struct {
		root              string
		m                 map[string]any
		createPlaceholder func(s string) string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "Test case 1",
			args: args{
				root: "root",
				m: map[string]any{
					"field name": "value",
				},
				createPlaceholder: func(s string) string {
					return "placeholder"
				},
			},
			want: map[string]any{
				"field name": "placeholder",
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			insertFieldPlaceholders(tt.args.root, tt.args.m, tt.args.createPlaceholder)
			if !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("insertFieldPlaceholders() = %v, want %v", tt.args.m, tt.want)
			}
		})
	}
}
