package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewComment_1(t *testing.T) {
	type args struct {
		pos  Pos
		text string
	}
	tests := []struct {
		name string
		args args
		want *CommentNode
	}{
		{
			name: "Test case 1",
			args: args{
				pos:  1,
				text: "This is a comment",
			},
			want: &CommentNode{
				NodeType: NodeComment,
				Pos:      1,
				tr:       nil,
				Text:     "This is a comment",
			},
		},
		{
			name: "Test case 2",
			args: args{
				pos:  2,
				text: "Another comment",
			},
			want: &CommentNode{
				NodeType: NodeComment,
				Pos:      2,
				tr:       nil,
				Text:     "Another comment",
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

			if got := (&Tree{}).newComment(tt.args.pos, tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
