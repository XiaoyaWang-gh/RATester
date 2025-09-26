package deploy

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestapplyOrdering_1(t *testing.T) {
	type args struct {
		ordering []*regexp.Regexp
		uploads  []*fileToUpload
	}
	tests := []struct {
		name string
		args args
		want [][]*fileToUpload
	}{
		{
			name: "Test Case 1",
			args: args{
				ordering: []*regexp.Regexp{
					regexp.MustCompile("^/path1/"),
					regexp.MustCompile("^/path2/"),
				},
				uploads: []*fileToUpload{
					{
						Local: &localFile{
							SlashPath: "/path1/file1",
						},
					},
					{
						Local: &localFile{
							SlashPath: "/path2/file2",
						},
					},
					{
						Local: &localFile{
							SlashPath: "/path3/file3",
						},
					},
				},
			},
			want: [][]*fileToUpload{
				{
					{
						Local: &localFile{
							SlashPath: "/path1/file1",
						},
					},
				},
				{
					{
						Local: &localFile{
							SlashPath: "/path2/file2",
						},
					},
				},
				{
					{
						Local: &localFile{
							SlashPath: "/path3/file3",
						},
					},
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

			if got := applyOrdering(tt.args.ordering, tt.args.uploads); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOrdering() = %v, want %v", got, tt.want)
			}
		})
	}
}
