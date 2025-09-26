package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMark_1(t *testing.T) {
	type testCase struct {
		name string
		s    *state
		want int
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			s: &state{
				vars: []variable{
					{name: "var1", value: reflect.ValueOf(1)},
					{name: "var2", value: reflect.ValueOf(2)},
					{name: "var3", value: reflect.ValueOf(3)},
				},
			},
			want: 3,
		},
		{
			name: "Test case 2",
			s: &state{
				vars: []variable{
					{name: "var1", value: reflect.ValueOf(1)},
					{name: "var2", value: reflect.ValueOf(2)},
				},
			},
			want: 2,
		},
		{
			name: "Test case 3",
			s: &state{
				vars: []variable{
					{name: "var1", value: reflect.ValueOf(1)},
				},
			},
			want: 1,
		},
		{
			name: "Test case 4",
			s: &state{
				vars: []variable{},
			},
			want: 0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.mark(); got != tt.want {
				t.Errorf("state.mark() = %v, want %v", got, tt.want)
			}
		})
	}
}
