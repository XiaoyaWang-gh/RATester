package constraints

import (
	"fmt"
	"testing"
)

func TestLabelFn_1(t *testing.T) {
	testCases := []struct {
		name   string
		value  string
		labels map[string]string
		want   bool
	}{
		{
			name:   "testLabel",
			value:  "testValue",
			labels: map[string]string{"testLabel": "testValue"},
			want:   true,
		},
		{
			name:   "testLabel",
			value:  "wrongValue",
			labels: map[string]string{"testLabel": "testValue"},
			want:   false,
		},
		{
			name:   "missingLabel",
			value:  "testValue",
			labels: map[string]string{"otherLabel": "testValue"},
			want:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fn := labelFn(tc.name, tc.value)
			got := fn(tc.labels)
			if got != tc.want {
				t.Errorf("labelFn(%s, %s) = %t, want %t", tc.name, tc.value, got, tc.want)
			}
		})
	}
}
