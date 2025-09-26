package fiber

import (
	"fmt"
	"testing"
)

func TestfindGreedyParamLen_1(t *testing.T) {
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
		{
			name: "Test case 1",
			args: args{
				s:           "abcdef",
				searchCount: 2,
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
				s:           "abcdefabcdef",
				searchCount: 2,
				segment: &routeSegment{
					ComparePart: "def",
					PartCount:   2,
				},
			},
			want: 6,
		},
		{
			name: "Test case 3",
			args: args{
				s:           "abcdefabcdef",
				searchCount: 3,
				segment: &routeSegment{
					ComparePart: "def",
					PartCount:   2,
				},
			},
			want: 0,
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
