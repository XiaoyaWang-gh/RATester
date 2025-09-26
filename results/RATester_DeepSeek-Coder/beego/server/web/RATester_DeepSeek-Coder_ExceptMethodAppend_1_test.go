package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExceptMethodAppend_1(t *testing.T) {
	testCases := []struct {
		name   string
		action string
		want   []string
	}{
		{
			name:   "Test case 1",
			action: "action1",
			want:   []string{"action1"},
		},
		{
			name:   "Test case 2",
			action: "action2",
			want:   []string{"action1", "action2"},
		},
		{
			name:   "Test case 3",
			action: "action3",
			want:   []string{"action1", "action2", "action3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			exceptMethod = []string{}
			ExceptMethodAppend(tc.action)
			if !reflect.DeepEqual(exceptMethod, tc.want) {
				t.Errorf("ExceptMethodAppend(%s) = %v, want %v", tc.action, exceptMethod, tc.want)
			}
		})
	}
}
