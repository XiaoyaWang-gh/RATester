package identity

import (
	"fmt"
	"testing"
)

func TestcheckMaxDepth_1(t *testing.T) {
	type args struct {
		sid   *searchID
		level int
	}
	tests := []struct {
		name string
		args args
		want FinderResult
	}{
		{
			name: "Test case 1",
			args: args{
				sid: &searchID{
					maxDepth: 10,
				},
				level: 5,
			},
			want: -1,
		},
		{
			name: "Test case 2",
			args: args{
				sid: &searchID{
					maxDepth: 10,
				},
				level: 11,
			},
			want: FinderNotFound,
		},
		{
			name: "Test case 3",
			args: args{
				sid: &searchID{
					maxDepth: -1,
				},
				level: 101,
			},
			want: FinderFound,
		},
		{
			name: "Test case 4",
			args: args{
				sid: &searchID{
					maxDepth: -1,
				},
				level: 100,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &Finder{}
			if got := f.checkMaxDepth(tt.args.sid, tt.args.level); got != tt.want {
				t.Errorf("Finder.checkMaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
