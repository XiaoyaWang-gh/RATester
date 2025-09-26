package fiber

import (
	"fmt"
	"testing"
)

func TestFindGreedyParamLen_1(t *testing.T) {
	type args struct {
		s           string
		searchCount int
		segment     *routeSegment
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
				s:           "abc/def/ghi",
				searchCount: 1,
				segment: &routeSegment{
					ComparePart: "def",
					PartCount:   1,
				},
			},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{
				s:           "abc/def/ghi",
				searchCount: 2,
				segment: &routeSegment{
					ComparePart: "def",
					PartCount:   2,
				},
			},
			want: 0,
		},
		{
			name: "Test case 3",
			args: args{
				s:           "abc/def/ghi",
				searchCount: 1,
				segment: &routeSegment{
					ComparePart: "ghi",
					PartCount:   1,
				},
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

			if got := findGreedyParamLen(tt.args.s, tt.args.searchCount, tt.args.segment); got != tt.want {
				t.Errorf("findGreedyParamLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
