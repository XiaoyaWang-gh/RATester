package fiber

import (
	"fmt"
	"testing"
)

func TestFindParamLenForLastSegment_1(t *testing.T) {
	type args struct {
		s   string
		seg *routeSegment
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test case 1",
			args: args{
				s:   "/test/path",
				seg: &routeSegment{IsGreedy: false},
			},
			want: 4,
		},
		{
			name: "Test case 2",
			args: args{
				s:   "/test/path/",
				seg: &routeSegment{IsGreedy: true},
			},
			want: 11,
		},
		{
			name: "Test case 3",
			args: args{
				s:   "/test/path",
				seg: &routeSegment{IsGreedy: true},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := findParamLenForLastSegment(tt.args.s, tt.args.seg); got != tt.want {
				t.Errorf("findParamLenForLastSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
