package rules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOrFunc_1(t *testing.T) {
	tests := []struct {
		name  string
		left  TreeBuilder
		right TreeBuilder
		want  *Tree
	}{
		{
			name: "Test case 1",
			left: func() *Tree {
				return &Tree{
					Matcher: "matcher1",
					Value:   []string{"value1", "value2"},
				}
			},
			right: func() *Tree {
				return &Tree{
					Matcher: "matcher2",
					Value:   []string{"value3", "value4"},
				}
			},
			want: &Tree{
				Matcher: "or",
				RuleLeft: &Tree{
					Matcher: "matcher1",
					Value:   []string{"value1", "value2"},
				},
				RuleRight: &Tree{
					Matcher: "matcher2",
					Value:   []string{"value3", "value4"},
				},
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

			got := orFunc(tt.left, tt.right)
			if !reflect.DeepEqual(got(), tt.want) {
				t.Errorf("orFunc() = %v, want %v", got(), tt.want)
			}
		})
	}
}
