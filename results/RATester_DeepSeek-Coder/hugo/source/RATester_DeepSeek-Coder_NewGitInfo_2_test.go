package source

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/bep/gitmap"
)

func TestNewGitInfo_2(t *testing.T) {
	type args struct {
		info gitmap.GitInfo
	}
	tests := []struct {
		name string
		args args
		want GitInfo
	}{
		{
			name: "Test case 1",
			args: args{
				info: gitmap.GitInfo{
					Hash:            "hash",
					AbbreviatedHash: "abbreviatedHash",
					Subject:         "subject",
					AuthorName:      "authorName",
					AuthorEmail:     "authorEmail",
					AuthorDate:      time.Now(),
					CommitDate:      time.Now(),
					Body:            "body",
				},
			},
			want: GitInfo{
				Hash:            "hash",
				AbbreviatedHash: "abbreviatedHash",
				Subject:         "subject",
				AuthorName:      "authorName",
				AuthorEmail:     "authorEmail",
				AuthorDate:      time.Now(),
				CommitDate:      time.Now(),
				Body:            "body",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewGitInfo(tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
