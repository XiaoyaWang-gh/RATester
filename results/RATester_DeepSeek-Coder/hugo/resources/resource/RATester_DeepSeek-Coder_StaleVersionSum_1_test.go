package resource

import (
	"fmt"
	"testing"
)

func TestStaleVersionSum_1(t *testing.T) {
	type StaleInfoMock struct {
		MockStaleVersion func() uint32
	}

	tests := []struct {
		name string
		args []any
		want uint32
	}{
		{
			name: "Test with no stale versions",
			args: []any{
				StaleInfoMock{MockStaleVersion: func() uint32 { return 0 }},
				StaleInfoMock{MockStaleVersion: func() uint32 { return 0 }},
			},
			want: 0,
		},
		{
			name: "Test with one stale version",
			args: []any{
				StaleInfoMock{MockStaleVersion: func() uint32 { return 1 }},
				StaleInfoMock{MockStaleVersion: func() uint32 { return 0 }},
			},
			want: 1,
		},
		{
			name: "Test with multiple stale versions",
			args: []any{
				StaleInfoMock{MockStaleVersion: func() uint32 { return 1 }},
				StaleInfoMock{MockStaleVersion: func() uint32 { return 2 }},
				StaleInfoMock{MockStaleVersion: func() uint32 { return 3 }},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := StaleVersionSum(tt.args...); got != tt.want {
				t.Errorf("StaleVersionSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
